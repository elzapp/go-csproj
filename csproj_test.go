package csproj

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	x := `<Project Sdk="Microsoft.NET.Sdk.Web">
	<PropertyGroup>
	  <TargetFramework>netcoreapp2.2</TargetFramework>
	  <NoWarn>$(NoWarn);CS1591</NoWarn>
	  <RootNamespace>Awesome.Elion</RootNamespace>
	  <GenerateDocumentationFile>true</GenerateDocumentationFile>
	  <PreserveCompilationContext>true</PreserveCompilationContext>
	  <AssemblyName>Awesome.Elion</AssemblyName>
	  <OutputType>Exe</OutputType>
	  <PackageId>Awesome.Elion</PackageId>
	</PropertyGroup>
	<PropertyGroup Condition="'$(Configuration)|$(Platform)'=='Debug|AnyCPU'">
	  <DocumentationFile>bin\Debug\netcoreapp2.1\Awesome.Elion.xml</DocumentationFile>
	</PropertyGroup>
	<ItemGroup>
	  <PackageReference Include="EPPlus" Version="4.5.3.2" />
	  <PackageReference Include="Jolly.Mestorf" Version="1.6.0" />
	  <PackageReference Include="Lonely.Jang" Version="2.1.0" />
	  <PackageReference Include="Fervent.Mayer" Version="1.4.2" />
	  <PackageReference Include="Romantic.Babbage" Version="1.0.0" />
	  <PackageReference Include="Dreamy.Kilby" Version="1.0.0" />
	  <PackageReference Include="Nauseous.Carson" Version="1.5.1" />
	  <PackageReference Include="Compassionate.Stonebraker" Version="2.4.0" />
	  <PackageReference Include="Jolly.Mahavira" Version="1.0.16" />
	  <PackageReference Include="Tiny.Gates" Version="1.0.14" />
	  <PackageReference Include="Backstabbing.Bardeen" Version="2.0.0" />
	  <PackageReference Include="Ecstatic.Bohr" Version="2.1.0" />
	  <PackageReference Include="Pedantic.Almeida" Version="1.0.0" />
	  <PackageReference Include="Focused.Jang" Version="0.1.1" />
	</ItemGroup>
  </Project>`
	fmt.Printf("%+v\n", ParseString(x))
}
