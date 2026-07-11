package ahoy

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func idGitRepo(t *testing.T, name, email string) string {
	t.Helper()
	t.Setenv("GIT_CONFIG_GLOBAL", os.DevNull)
	t.Setenv("GIT_CONFIG_SYSTEM", os.DevNull)
	dir := t.TempDir()
	idMustGit(t, dir, "init")
	if name != "" {
		idMustGit(t, dir, "config", "user.name", name)
	}
	if email != "" {
		idMustGit(t, dir, "config", "user.email", email)
	}
	return dir
}

func idMustGit(t *testing.T, dir string, args ...string) {
	t.Helper()
	if out, err := exec.Command("git", append([]string{"-C", dir}, args...)...).CombinedOutput(); err != nil {
		t.Fatalf("git %v: %v\n%s", args, err, out)
	}
}

func idWritePin(t *testing.T, root, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Join(root, ".abcd", "config"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, ".abcd", "config", "identity.json"), []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestDetectGitIdentity_Mismatch(t *testing.T) {
	dir := idGitRepo(t, "Test User", "test@example.com")
	idWritePin(t, dir, `{"name":"Alex Reppel","email":"alex@example.com"}`)
	gaps := detectGitIdentity(dir)
	if len(gaps) != 1 || gaps[0].ID != "git_identity.mismatch" || !gaps[0].Required {
		t.Fatalf("want one required git_identity.mismatch gap, got %+v", gaps)
	}
}

func TestDetectGitIdentity_Match(t *testing.T) {
	dir := idGitRepo(t, "Alex Reppel", "alex@example.com")
	idWritePin(t, dir, `{"name":"Alex Reppel","email":"alex@example.com"}`)
	if gaps := detectGitIdentity(dir); len(gaps) != 0 {
		t.Fatalf("want no gap on match, got %+v", gaps)
	}
}

func TestDetectGitIdentity_Unpinned(t *testing.T) {
	dir := idGitRepo(t, "Alex Reppel", "alex@example.com")
	gaps := detectGitIdentity(dir)
	if len(gaps) != 1 || gaps[0].ID != "git_identity.unpinned" || gaps[0].Required {
		t.Fatalf("want one advisory git_identity.unpinned gap, got %+v", gaps)
	}
}
