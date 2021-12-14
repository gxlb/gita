package gitlib

import (
	"fmt"
	"gita/exec"
	"regexp"
	"strings"
)

type BranchTreeInfo struct {
	Branch string
	Prefix string
	Depth  int
}

func (bi *BranchTreeInfo) ValidateSync(name string) error {
	if bi == nil {
		return fmt.Errorf("invalid branch '%s'", name)
	}
	if bi.Prefix != topTree {
		return fmt.Errorf("branch '%s' not under top tree", bi.Branch)
	}
	return nil
}

type BranchesGraphInfo struct {
	Local       map[string]*BranchTreeInfo
	Remote      map[string]*BranchTreeInfo
	Current     string
	CurrentInfo BranchTreeInfo
}

func (gi *BranchesGraphInfo) GetBranchInfo(branch string, remote bool) *BranchTreeInfo {
	if remote {
		return gi.Remote[branch]
	}
	return gi.Local[branch]
}

func (gi *BranchesGraphInfo) CheckSyncBranch(branch string) (*BranchTreeInfo, error) {
	b := gi.GetBranchInfo(branch, true)
	if err := b.ValidateSync(remoteMapPrefix + branch); err != nil {
		return nil, err
	}
	if local := gi.GetBranchInfo(branch, false); local != nil { // check only if local branch exists
		if err := local.ValidateSync(branch); err != nil {
			return nil, err
		}
		if local.Depth < b.Depth {
			return nil, fmt.Errorf("local branch '%s' is ahead of remote", branch)
		}
	}
	return b, nil
}

func (gi *BranchesGraphInfo) ValidateSync(src *BranchTreeInfo, dst string) (needSync bool, err error) {
	di, err := gi.CheckSyncBranch(dst)
	if err != nil {
		return false, err
	}
	if di.Branch == src.Branch {
		return false, fmt.Errorf("dst branch '%s' is the src branch", dst)
	}
	if di.Depth < src.Depth {
		return false, fmt.Errorf("dst branch '%s' is ahead of src branch '%s'",
			di.Branch, src.Branch)
	}
	needSync = src.Depth < di.Depth
	return needSync, nil
}

func (gi *BranchesGraphInfo) ValidateSyncList(src string, dsts []string) ([]string, error) {
	if err := ValidateCheckout(); err != nil {
		return nil, err
	}
	si, err := gi.CheckSyncBranch(src)
	if err != nil {
		return nil, err
	}
	ret := dsts[:0]
	for _, dst := range dsts {
		needSync, err := gi.ValidateSync(si, dst)
		if err != nil {
			return nil, err
		}
		if needSync {
			ret = append(ret, dst)
		}
	}
	return ret, nil
}

func (gi *BranchesGraphInfo) SyncBranchList(src string, dsts []string, realDoSync bool) error {
	needSync, err := gi.ValidateSyncList(src, dsts)
	if err != nil {
		return err
	}

	if len(needSync) > 0 {
		for _, dst := range needSync {
			if err := gi.syncBranch(src, dst, realDoSync); err != nil {
				return err
			}
		}
		if realDoSync {
			if err := Checkout(gi.Current, false, false); err != nil {
				return err
			}
		}
	}
	return nil
}

func (gi *BranchesGraphInfo) syncBranch(src string, dst string, realDoSync bool) error {
	cmds := []string{
		fmt.Sprintf(CmdForceCheckoutRemoteF, dst, dst),
		fmt.Sprintf(`git pull --progress -v --no-rebase origin %s`, src),
		fmt.Sprintf(`git push --progress origin %s:%s`, dst, dst),
	}
	if realDoSync {
		if err := exec.ShellExecuteList(cmds); err != nil {
			return err
		}
	} else {
		PrettyShow(cmds)
	}

	return nil
}

/*
	* cbc1120 (origin/develop, origin/dev_demo, origin/alpha, dev_demo) style: add todo dismiss pathType warning log
	| * 9ad90a3 (develop) style: update db schema
	|/
	* cdc1bd1 (origin/feature/service) feat: any method supported on polyapi
	* 6c39422 (origin/release/v0.7.2, origin/fix/reqPoly, origin/dev_nopoly) fix: item always nil
	| * 0455241 (origin/release/v0.7.0_nopoly) feat: remove poly api feature
	|/
	| * f5eba4a (origin/release/info-ms) fix:raw_api add req info
	|/
	* 4ba7228 (tag: v0.7, origin/release/v0.7.0, origin/master) fix: verify signcmd when create service
	* b9916cd (origin/fix/upload) fix: get params from form-data
	| * c0b01ad (origin/patch-1) Update .gitlab-ci.yml
	|/
	* 9a8dbf8 (alpha) fix: query & formData parameter in swagger inputs
	* 5822a11 (feature/service) fix: update db type
	| * 953c81d (HEAD -> master, tag: v0.6.3, origin/release/v0.6.1, origin/fix/v0.6.1-r2, release/v0.6.1) fix: log & response of raw api request
	|/
	* 8fe25cb (origin/fix/v0.6.1-r1, fix/v0.6.1-r1) style: update data fix sql
	* 21bbf4b (tag: v0.6.1) fix: golint error
	| * 7593c1d (origin/bak/master0.0.1, bak/master0.0.1) fix: delete unused files
	| *   17334b0 Merge branch 'release' of https://git.internal.yunify.com/qxp/polyapi
	| |\
	| |/
	|/|
	* | c13c986 (tag: v0.6.0.R1) fix: expose port 9090 in Dockerfile
	| * 19dca67 (tag: v0.0.1) Merge branch 'debug' into 'master'
	|/
	* 51483f1 upload README
*/
const exprBranchMap = `(?m-s:^\s*(?P<PRE>[^0-9a-f\n]+)\s+[0-9a-f]+\s+\((?P<BRANCHES>[^\n\)]*?)\)[^\n]*$)`

var regexpBranchMap = regexp.MustCompile(exprBranchMap)

const (
	headMapPrefix   = "HEAD -> "
	tagMapPrefix    = "tag: "
	remoteMapPrefix = "origin/"
	topTree         = "*"
)

var (
	exprBranchPrefix = fmt.Sprintf(`(?m-s:^\s?(?P<PREFIX>%s|%s|%s|\s?)(?P<BRANCH>[-\w\./]+)$)`,
		headMapPrefix, tagMapPrefix, remoteMapPrefix)
	regexpBranchPrefix = regexp.MustCompile(exprBranchPrefix)
)

func GetBranchesGraph() (*BranchesGraphInfo, error) {
	Fetch()
	r, err := exec.ShellExecuteResult(CmdBranchesGraph)
	if err != nil {
		return nil, err
	}

	ret := &BranchesGraphInfo{
		Local:  map[string]*BranchTreeInfo{},
		Remote: map[string]*BranchTreeInfo{},
	}
	treeDepth := -1
	regexpBranchMap.ReplaceAllStringFunc(r, func(src string) string {
		treeDepth++
		elems := regexpBranchMap.FindAllStringSubmatch(src, 1)[0]
		prefix, branches := elems[1], elems[2]
		if prefix != topTree && strings.HasPrefix(prefix, topTree) {
			fmt.Printf("@@ reset tree prefix '%s' => '%s' (%s)\n", prefix, topTree, branches)
			prefix = topTree
		}
		bs := strings.Split(branches, ",")
		for _, v := range bs {
			elems := regexpBranchPrefix.FindAllStringSubmatch(v, 1)
			if len(elems) > 0 {
				pre, branch := elems[0][1], elems[0][2]
				switch pre {
				case headMapPrefix:
					ret.Current = branch
					ret.CurrentInfo.Prefix = prefix
					ret.CurrentInfo.Depth = treeDepth
					ret.CurrentInfo.Branch = branch
					fallthrough
				default:
					ret.Local[branch] = &BranchTreeInfo{
						Prefix: prefix,
						Depth:  treeDepth,
						Branch: branch,
					}
				case remoteMapPrefix:
					ret.Remote[branch] = &BranchTreeInfo{
						Prefix: prefix,
						Depth:  treeDepth,
						Branch: remoteMapPrefix + branch,
					}
				case tagMapPrefix:
					// do nothing
				}
			} else {
				fmt.Printf("**  mismatch branch pattern: %q\n", v)
			}
		}
		return src
	})
	return ret, nil
}

func ShowBranchesGraph() error {
	return exec.ShellExecute(CmdBranchesGraph)
}

func ShowBranches(all bool) error {
	return nil
}
