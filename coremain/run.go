// Modifications by Lightstream: Auto-updating, built-in configuration and logging.

// Package coremain contains the functions for starting CoreDNS.
package coremain

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"net"

	"github.com/blang/semver"
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

func doSelfUpdate() {
	v := semver.MustParse(appVersion)
	latest, err := selfupdate.UpdateSelf(v, "golightstream/lightstream-prism-dns")
	if err != nil {
		log.Println("Binary update failed:", err)
		return
	}
	if latest.Version.LTE(v) {
		// latest version is the same as current version. It means current binary is up to date.
		if latest.Version.Equals(v) {
			log.Println("Current binary is the latest version", v)
		} else {
			log.Printf("You have an unreleased/newer version than latest; %v > latest (%v)\n", v, latest.Version)
		}
	} else {
		log.Println("Successfully updated to version", latest.Version)
		log.Println("Release note:\n", latest.ReleaseNotes)
	}
}

func init() {
	caddy.DefaultConfigFile = "Corefile"
	caddy.Quiet = true // don't show init stuff from caddy
	setVersion()

	flag.StringVar(&conf, "conf", "", "Corefile to load (default \""+caddy.DefaultConfigFile+"\")")
	flag.BoolVar(&plugins, "plugins", false, "List installed plugins")
	flag.StringVar(&caddy.PidFile, "pidfile", "", "Path to write pid file")
	flag.BoolVar(&version, "version", false, "Show version")
	flag.BoolVar(&dnsserver.Quiet, "quiet", false, "Quiet mode (no initialization output)")

	caddy.RegisterCaddyfileLoader("flag", caddy.LoaderFunc(confLoader))
	caddy.SetDefaultCaddyfileLoader("default", caddy.LoaderFunc(defaultLoader))

	flag.StringVar(&dnsserver.Port, serverType+".port", dnsserver.DefaultPort, "Default port")
	flag.StringVar(&dnsserver.Port, "p", dnsserver.DefaultPort, "Default port")

	caddy.AppName = coreName
	caddy.AppVersion = CoreVersion
}

// Run is CoreDNS's main() function.
func Run() {
	caddy.TrapSignals()
	flag.Parse()

	if len(flag.Args()) > 0 {
		mustLogFatal(fmt.Errorf("extra command line arguments: %s", flag.Args()))
	}

	log.SetOutput(os.Stdout)
	log.SetFlags(0) // Set to 0 because we're doing our own time, with timezone

	if version {
		showVersion()
		os.Exit(0)
	}
	if plugins {
		fmt.Println(caddy.DescribePlugins())
		os.Exit(0)
	}

	// Get Corefile input
	corefile, err := caddy.LoadCaddyfile(serverType)
	if err != nil {
		mustLogFatal(err)
	}

	ifaces, _ := net.Interfaces()

	log.Printf(
		" _\n" +
			"| |   (_) __ _| |__ | |_ ___| |_ _ __ ___  __ _ _ __ ___\n" +
			"| |   | |/ _``| '_ \\| __/ __| __| '__/ _ \\/ _` | '_ ` _ \\\n" +
			"| |___| | (_| | | | | |_\\__ \\ |_| | |  __/ (_| | | | | | |\n" +
			"|_____|_|\\__, |_| |_|\\__|___/\\__|_|  \\___|\\__,_|_| |_| |_|\n" +
			"         |___/")

	doSelfUpdate()

	log.Printf("Lightstream Console DNS listening on:")
	for _, i := range ifaces {
		if addrs, err := i.Addrs(); err == nil {
			for _, addr := range addrs {
				switch v := addr.(type) {
				case *net.IPNet:
					if v.IP.To4() != nil {
						if v.IP.String() != "127.0.0.1" {
							log.Printf("-> %s", v.IP)
						}

					}
				}
			}
		}
	}

	log.Printf("If listening fails with permission denied, please run in privileged mode (Admin/Root)")

	// Start your engines
	instance, err := caddy.Start(corefile)
	if err != nil {
		mustLogFatal(err)
	}

	// if !dnsserver.Quiet {
	// 	// showVersion()
	// }

	// Twiddle your thumbs
	instance.Wait()
}

// mustLogFatal wraps log.Fatal() in a way that ensures the
// output is always printed to stderr so the user can see it
// if the user is still there, even if the process log was not
// enabled. If this process is an upgrade, however, and the user
// might not be there anymore, this just logs to the process
// log and exits.
func mustLogFatal(args ...interface{}) {
	if !caddy.IsUpgrade() {
		log.SetOutput(os.Stderr)
	}
	log.Fatal(args...)
}

// confLoader loads the Caddyfile using the -conf flag.
func confLoader(serverType string) (caddy.Input, error) {
	if conf == "" {
		defaultConfString := `
		. {
			forward . 1.1.1.1 8.8.8.8
			dns64 {
        		allow_ipv4
    		}
			rewrite continue {
				name regex live(.*).twitch.tv live{1}.int01.golightstream.com
				answer name live(.*).int01.golightstream.com live{1}.twitch.tv
			}
			rewrite continue {
				name regex (.*).contribute.live-video.net live{1}.int01.golightstream.com
				answer name live(.*).int01.golightstream.com {1}.contribute.live-video.net
			}
			rewrite continue {
				name regex (.*).psdnstest.golightstream.com d3exjhue0wekgd.cloudfront.net answer auto
			}
			log . {
				class all
			}
	}`
		return caddy.CaddyfileInput{
			Contents:       []byte(defaultConfString),
			Filepath:       conf,
			ServerTypeName: serverType,
		}, nil
	}

	if conf == "stdin" {
		return caddy.CaddyfileFromPipe(os.Stdin, serverType)
	}

	contents, err := os.ReadFile(filepath.Clean(conf))
	if err != nil {
		return nil, err
	}
	return caddy.CaddyfileInput{
		Contents:       contents,
		Filepath:       conf,
		ServerTypeName: serverType,
	}, nil
}

// defaultLoader loads the Corefile from the current working directory.
func defaultLoader(serverType string) (caddy.Input, error) {
	contents, err := os.ReadFile(caddy.DefaultConfigFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	return caddy.CaddyfileInput{
		Contents:       contents,
		Filepath:       caddy.DefaultConfigFile,
		ServerTypeName: serverType,
	}, nil
}

// showVersion prints the version that is starting.
func showVersion() {
	fmt.Print(versionString())
	fmt.Print(releaseString())
	if devBuild && gitShortStat != "" {
		fmt.Printf("%s\n%s\n", gitShortStat, gitFilesModified)
	}
}

// versionString returns the CoreDNS version as a string.
func versionString() string {
	return fmt.Sprintf("%s-%s\n", caddy.AppName, caddy.AppVersion)
}

// releaseString returns the release information related to CoreDNS version:
// <OS>/<ARCH>, <go version>, <commit>
// e.g.,
// linux/amd64, go1.8.3, a6d2d7b5
func releaseString() string {
	return fmt.Sprintf("%s/%s, %s, %s\n", runtime.GOOS, runtime.GOARCH, runtime.Version(), GitCommit)
}

// setVersion figures out the version information
// based on variables set by -ldflags.
func setVersion() {
	gitTag = GitCommit

	// A development build is one that's not at a tag or has uncommitted changes
	devBuild = gitTag == "" || gitShortStat != ""

	// Only set the appVersion if -ldflags was used
	if gitNearestTag != "" || gitTag != "" {
		if devBuild && gitNearestTag != "" {
			appVersion = fmt.Sprintf("%s (+%s %s)", strings.TrimPrefix(gitNearestTag, "v"), GitCommit, buildDate)
		} else if gitTag != "" {
			appVersion = strings.TrimPrefix(gitTag, "v")
		}
	}
}

// Flags that control program flow or startup
var (
	conf    string
	version bool
	plugins bool
)

// Build information obtained with the help of -ldflags
var (
	// nolint
	appVersion = "(untracked dev build)" // inferred at startup
	devBuild   = true                    // inferred at startup

	buildDate        string // date -u
	gitTag           string // git describe --exact-match HEAD 2> /dev/null
	gitNearestTag    string // git describe --abbrev=0 --tags HEAD
	gitShortStat     string // git diff-index --shortstat
	gitFilesModified string // git diff-index --name-only HEAD

	// Gitcommit contains the commit where we built CoreDNS from.
	GitCommit string
)
