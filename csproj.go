package csproj

import "encoding/xml"

type Project struct {
	TargetFramework                     string       `xml:"PropertyGroup>TargetFramework"`
	NoWarn                              string       `xml:"PropertyGroup>NoWarn"`
	GenerateDocumentationFile           string       `xml:"PropertyGroup>GenerateDocumentationFile"`
	PreserveCompilationContext          string       `xml:"PropertyGroup>PreserveCompilationContext"`
	AssemblyName                        string       `xml:"PropertyGroup>AssemblyName"`
	OutputType                          string       `xml:"PropertyGroup>OutputType"`
	PackageID                           string       `xml:"PropertyGroup>PackageId"`
	RootNamespace                       string       `xml:"PropertyGroup>RootNamespace"`
	PublishWithAspNetCoreTargetManifest string       `xml:"PropertyGroup>PublishWithAspNetCoreTargetManifest"`
	UserSecretsID                       string       `xml:"PropertyGroup>UserSecretsId"`
	Version                             string       `xml:"PropertyGroup>Version"`
	Dependencies                        []Dependency `xml:"ItemGroup>PackageReference"`
	AspNetCoreHostingModel              string       `xml:"PropertyGroup>AspNetCoreHostingModel"`
	DockerDefaultTargetOS               string       `xml:"PropertyGroup>DockerDefaultTargetOS"`
}

type Dependency struct {
	Version   string `xml:"Version,attr"`
	PackageID string `xml:"Include,attr"`
}

func ParseString(doc string) (p Project) {
	xml.Unmarshal([]byte(doc), &p)
	return
}
