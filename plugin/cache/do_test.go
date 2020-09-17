package cache

import (
	"context"
	"testing"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/pkg/dnstest"
	"github.com/coredns/coredns/plugin/test"

	"github.com/miekg/dns"
)

func TestDo(t *testing.T) {
	// cache sets Do and requests that don't have them.
	c := New()
	c.Next = echoHandler()
	req := new(dns.Msg)
	req.SetQuestion("example.org.", dns.TypeA)
	rec := dnstest.NewRecorder(&test.ResponseWriter{})

	// No DO set.
	c.ServeDNS(context.TODO(), rec, req)
	reply := rec.Msg
	opt := reply.Extra[len(reply.Extra)-1]
	if x, ok := opt.(*dns.OPT); !ok {
		t.Fatalf("Expected OPT RR, got %T", x)
	}
	if !opt.(*dns.OPT).Do() {
		t.Errorf("Expected DO bit to be set, got false")
	}
	if x := opt.(*dns.OPT).UDPSize(); x != defaultUDPBufSize {
		t.Errorf("Expected %d bufsize, got %d", defaultUDPBufSize, x)
	}

	// Do set - so left alone.
	const mysize = defaultUDPBufSize * 2
	setDo(req)
	// set bufsize to something else than default to see cache doesn't touch it
	req.Extra[len(req.Extra)-1].(*dns.OPT).SetUDPSize(mysize)
	c.ServeDNS(context.TODO(), rec, req)
	reply = rec.Msg
	opt = reply.Extra[len(reply.Extra)-1]
	if x, ok := opt.(*dns.OPT); !ok {
		t.Fatalf("Expected OPT RR, got %T", x)
	}
	if !opt.(*dns.OPT).Do() {
		t.Errorf("Expected DO bit to be set, got false")
	}
	if x := opt.(*dns.OPT).UDPSize(); x != mysize {
		t.Errorf("Expected %d bufsize, got %d", mysize, x)
	}

	// edns0 set, but not DO, so _not_ left alone.
	req.Extra[len(req.Extra)-1].(*dns.OPT).SetDo(false)
	c.ServeDNS(context.TODO(), rec, req)
	reply = rec.Msg
	opt = reply.Extra[len(reply.Extra)-1]
	if x, ok := opt.(*dns.OPT); !ok {
		t.Fatalf("Expected OPT RR, got %T", x)
	}
	if !opt.(*dns.OPT).Do() {
		t.Errorf("Expected DO bit to be set, got false")
	}
	if x := opt.(*dns.OPT).UDPSize(); x != defaultUDPBufSize {
		t.Errorf("Expected %d bufsize, got %d", defaultUDPBufSize, x)
	}
}

func echoHandler() plugin.Handler {
	return plugin.HandlerFunc(func(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
		w.WriteMsg(r)
		return dns.RcodeSuccess, nil
	})
}
