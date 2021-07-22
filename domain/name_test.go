package domain_test

import (
	"testing"

	"github.com/detectify/n5/domain"
	"github.com/stretchr/testify/require"
)

func TestDomain_WithSimpleName_ShouldParse(t *testing.T) {
	name, err := domain.Parse("com")
	require.NoError(t, err)
	require.Equal(t, "com", name.String())
	require.Equal(t, "com.", name.FQDN())
	require.Equal(t, "", name.Apex().String())
	require.Equal(t, "", name.Subdomain())
	require.Equal(t, "com", name.EffectiveTLD())

	name, err = domain.Parse("foo.com")
	require.NoError(t, err)
	require.Equal(t, "foo.com", name.String())
	require.Equal(t, "foo.com.", name.FQDN())
	require.Equal(t, "foo.com", name.Apex().String())
	require.Equal(t, "", name.Subdomain())
	require.Equal(t, "com", name.EffectiveTLD())

	name, err = domain.Parse("foo.bar.com")
	require.NoError(t, err)
	require.Equal(t, "foo.bar.com", name.String())
	require.Equal(t, "foo.bar.com.", name.FQDN())
	require.Equal(t, "bar.com", name.Apex().String())
	require.Equal(t, "foo", name.Subdomain())
	require.Equal(t, "com", name.EffectiveTLD())

	name, err = domain.Parse("foo.bar.baz.com")
	require.NoError(t, err)
	require.Equal(t, "foo.bar.baz.com", name.String())
	require.Equal(t, "baz.com", name.Apex().String())
	require.Equal(t, "foo.bar", name.Subdomain())
	require.Equal(t, "com", name.EffectiveTLD())

	name, err = domain.Parse("foo.bar.baz.com.")
	require.NoError(t, err)
	require.Equal(t, "foo.bar.baz.com", name.String())

	name, err = domain.Parse("foo.b_a_r.baz.com")
	require.NoError(t, err)
	require.Equal(t, "foo.b_a_r.baz.com", name.String())

	name, err = domain.Parse("foo.b-a-r.baz.com")
	require.NoError(t, err)
	require.Equal(t, "foo.b-a-r.baz.com", name.String())

	name, err = domain.Parse("www.ck")
	require.NoError(t, err)
	require.Equal(t, "www.ck", name.String())
	require.Equal(t, "www.ck", name.Apex().String())
	require.Equal(t, "", name.Subdomain())
	require.Equal(t, "ck", name.EffectiveTLD())
}

func TestDomain_WithNonIDNName_ShouldParse(t *testing.T) {
	name, err := domain.Parse("exåmple.com")
	require.NoError(t, err)
	require.Equal(t, "xn--exmple-jua.com", name.String())
	require.Equal(t, "exåmple.com", name.Unicode())
	require.True(t, name.HasPublicSuffix())
	require.True(t, name.IsICANN())

	name, err = domain.Parse("www.ēxämplĕ.cöm")
	require.NoError(t, err)
	require.Equal(t, "www.xn--xmpl-loa4zta.xn--cm-fka", name.String())
	require.Equal(t, "xn--xmpl-loa4zta.xn--cm-fka", name.Apex().String())
	require.Equal(t, "www", name.Subdomain())
	require.Equal(t, "xn--cm-fka", name.EffectiveTLD())
	require.Equal(t, "www.ēxämplĕ.cöm", name.Unicode())
	require.False(t, name.HasPublicSuffix())
	require.False(t, name.IsICANN())

	name, err = domain.Parse("пример.мкд")
	require.NoError(t, err)
	require.Equal(t, "xn--e1afmkfd.xn--d1alf", name.String())
	require.Equal(t, "xn--e1afmkfd.xn--d1alf", name.Apex().String())
	require.Equal(t, "", name.Subdomain())
	require.Equal(t, "xn--d1alf", name.EffectiveTLD())
	require.Equal(t, "пример.мкд", name.Unicode())
	require.True(t, name.HasPublicSuffix())
	require.True(t, name.IsICANN())
}

func TestDomain_WithMultiLabelETLD_ShouldParse(t *testing.T) {
	name, err := domain.Parse("foo.bar.baz.co.uk")

	require.NoError(t, err)
	require.Equal(t, "foo.bar.baz.co.uk", name.String())
	require.Equal(t, "baz.co.uk", name.Apex().String())
	require.Equal(t, "foo.bar", name.Subdomain())
	require.Equal(t, "co.uk", name.EffectiveTLD())
	require.True(t, name.IsICANN())

	name, err = domain.Parse("ec2-52-201-222-125.compute-1.amazonaws.com")
	require.NoError(t, err)
	require.Equal(t, "ec2-52-201-222-125.compute-1.amazonaws.com", name.String())
	require.Equal(t, "", name.Apex().String())
	require.Equal(t, "", name.Subdomain())
	require.Equal(t, "ec2-52-201-222-125.compute-1.amazonaws.com", name.EffectiveTLD())

	name, err = domain.Parse("foo.bar.blogspot.co.ke")
	require.NoError(t, err)
	require.Equal(t, "foo.bar.blogspot.co.ke", name.String())
	require.Equal(t, "bar.blogspot.co.ke", name.Apex().String())
	require.Equal(t, "foo", name.Subdomain())
	require.Equal(t, "blogspot.co.ke", name.EffectiveTLD())

	name, err = domain.Parse("www.accident-investigation.aero")
	require.NoError(t, err)
	require.Equal(t, "www.accident-investigation.aero", name.String())
	require.Equal(t, "www.accident-investigation.aero", name.Apex().String())
	require.Equal(t, "", name.Subdomain())
	require.Equal(t, "accident-investigation.aero", name.EffectiveTLD())

	name, err = domain.Parse("city.kobe.jp")
	require.NoError(t, err)
	require.Equal(t, "city.kobe.jp", name.String())
	require.Equal(t, "city.kobe.jp", name.Apex().String())
	require.Equal(t, "", name.Subdomain())
	require.Equal(t, "kobe.jp", name.EffectiveTLD())

	name, err = domain.Parse("www.city.kobe.jp")
	require.NoError(t, err)
	require.Equal(t, "www.city.kobe.jp", name.String())
	require.Equal(t, "city.kobe.jp", name.Apex().String())
	require.Equal(t, "www", name.Subdomain())
	require.Equal(t, "kobe.jp", name.EffectiveTLD())

	name, err = domain.Parse("www.fisk.kobe.jp")
	require.NoError(t, err)
	require.Equal(t, "www.fisk.kobe.jp", name.String())
	require.Equal(t, "www.fisk.kobe.jp", name.Apex().String())
	require.Equal(t, "", name.Subdomain())
	require.Equal(t, "fisk.kobe.jp", name.EffectiveTLD())

	name, err = domain.Parse("www.mm.kobe.jp")
	require.NoError(t, err)
	require.Equal(t, "www.mm.kobe.jp", name.String())
	require.Equal(t, "www.mm.kobe.jp", name.Apex().String())
	require.Equal(t, "", name.Subdomain())
	require.Equal(t, "mm.kobe.jp", name.EffectiveTLD())

	name, err = domain.Parse("www.coop.km")
	require.NoError(t, err)
	require.Equal(t, "www.coop.km", name.String())
	require.Equal(t, "www.coop.km", name.Apex().String())
	require.Equal(t, "", name.Subdomain())
	require.Equal(t, "coop.km", name.EffectiveTLD())

	name, err = domain.Parse("foo.web.tr")
	require.NoError(t, err)
	require.Equal(t, "foo.web.tr", name.String())
	require.Equal(t, "foo.web.tr", name.Apex().String())
	require.Equal(t, "", name.Subdomain())
	require.Equal(t, "web.tr", name.EffectiveTLD())

	name, err = domain.Parse("foo.alwaysdata.net")
	require.NoError(t, err)
	require.Equal(t, "foo.alwaysdata.net", name.String())
	require.Equal(t, "foo.alwaysdata.net", name.Apex().String())
	require.Equal(t, "", name.Subdomain())
	require.Equal(t, "alwaysdata.net", name.EffectiveTLD())
}

func TestDomain_WithNonICANNName_ShouldParse(t *testing.T) {
	name, err := domain.Parse("foo.bar.baz")
	require.NoError(t, err)
	require.Equal(t, "foo.bar.baz", name.String())
	require.Equal(t, "bar.baz", name.Apex().String())
	require.Equal(t, "foo", name.Subdomain())
	require.Equal(t, "baz", name.EffectiveTLD())
	require.False(t, name.IsICANN())
	require.False(t, name.HasPublicSuffix())

	name, err = domain.Parse("ec2-52-201-222-125.compute-1.amazonaws.com")
	require.NoError(t, err)
	require.Equal(t, "ec2-52-201-222-125.compute-1.amazonaws.com", name.String())
	require.Equal(t, "", name.Apex().String())
	require.Equal(t, "", name.Subdomain())
	require.Equal(t, "ec2-52-201-222-125.compute-1.amazonaws.com", name.EffectiveTLD())
	require.False(t, name.IsICANN())
	require.False(t, name.HasPublicSuffix())

	name, err = domain.Parse("foo.bar.ec2-52-201-222-125.compute-1.amazonaws.com")
	require.NoError(t, err)
	require.Equal(t, "foo.bar.ec2-52-201-222-125.compute-1.amazonaws.com", name.String())
	require.Equal(t, "bar.ec2-52-201-222-125.compute-1.amazonaws.com", name.Apex().String())
	require.Equal(t, "foo", name.Subdomain())
	require.Equal(t, "ec2-52-201-222-125.compute-1.amazonaws.com", name.EffectiveTLD())
	require.False(t, name.IsICANN())
	require.True(t, name.HasPublicSuffix())

	name, err = domain.Parse("www.geek")
	require.NoError(t, err)
	require.Equal(t, "www.geek", name.String())
	require.Equal(t, "www.geek", name.Apex().String())
	require.Equal(t, "", name.Subdomain())
	require.Equal(t, "geek", name.EffectiveTLD())
	require.False(t, name.IsICANN())
	require.False(t, name.HasPublicSuffix())
}

func TestDomain_WithMissingLabels_ShouldNotParse(t *testing.T) {
	_, err := domain.Parse(".")
	require.Error(t, err)

	_, err = domain.Parse("err..foo.com")
	require.Error(t, err)

	_, err = domain.Parse("www..uk")
	require.Error(t, err)
}

func TestDomain_WithMalformedLabels_ShouldNotParse(t *testing.T) {
	_, err := domain.Parse("@foo.com")
	require.Error(t, err)

	_, err = domain.Parse("foo-.com")
	require.Error(t, err)

	_, err = domain.Parse("-foo.com")
	require.Error(t, err)

	_, err = domain.Parse("foo!.com")
	require.Error(t, err)

	_, err = domain.Parse("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa.com")
	require.Error(t, err)

	_, err = domain.Parse("fo?o.com")
	require.Error(t, err)
}

func TestDomain_WithWildcard_ShouldParse(t *testing.T) {
	name, err := domain.Parse("*.example.com")
	require.NoError(t, err)
	require.Equal(t, "example.com", name.String())
	require.Equal(t, "example.com", name.Apex().String())
	require.Equal(t, "", name.Subdomain())
	require.Equal(t, "com", name.EffectiveTLD())

	name, err = domain.Parse("@.blog.example.com")
	require.NoError(t, err)
	require.Equal(t, "blog.example.com", name.String())
	require.Equal(t, "example.com", name.Apex().String())
	require.Equal(t, "blog", name.Subdomain())
	require.Equal(t, "com", name.EffectiveTLD())
}

func TestDomain_WithFQDN_ShouldParse(t *testing.T) {
	name, err := domain.Parse("example.com.")
	require.NoError(t, err)
	require.Equal(t, "example.com", name.String())

	name, err = domain.Parse("*.example.com.")
	require.NoError(t, err)
	require.Equal(t, "example.com", name.String())

	name, err = domain.Parse("sub.example.com.")
	require.NoError(t, err)
	require.Equal(t, "sub.example.com", name.String())

	name, err = domain.Parse("com.")
	require.NoError(t, err)
	require.Equal(t, "com", name.String())
}
