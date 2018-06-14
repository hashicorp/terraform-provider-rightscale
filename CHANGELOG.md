# Change Log

## [Unreleased](https://github.com/rightscale/terraform-provider-rightscale/tree/HEAD)

[Full Changelog](https://github.com/rightscale/terraform-provider-rightscale/compare/v1.1.0...HEAD)

**Closed issues:**

- Terraform Provider Development Program - Initial Review [\#60](https://github.com/rightscale/terraform-provider-rightscale/issues/60)

**Merged pull requests:**

- Steel 260 provider hash acc modifications [\#61](https://github.com/rightscale/terraform-provider-rightscale/pull/61) ([mastamark](https://github.com/mastamark))
- STEEL-290 cwf array type [\#59](https://github.com/rightscale/terraform-provider-rightscale/pull/59) ([crunis](https://github.com/crunis))
- STEEL-284 CWF Documentation fix [\#58](https://github.com/rightscale/terraform-provider-rightscale/pull/58) ([crunis](https://github.com/crunis))
- STEEL-281 Fix server delete [\#57](https://github.com/rightscale/terraform-provider-rightscale/pull/57) ([crunis](https://github.com/crunis))
- STEEL-277 Clean up code [\#56](https://github.com/rightscale/terraform-provider-rightscale/pull/56) ([crunis](https://github.com/crunis))
- STEEL-283 - 'private\_ip\_address' for instance now added as a field so… [\#55](https://github.com/rightscale/terraform-provider-rightscale/pull/55) ([mastamark](https://github.com/mastamark))
- STEEL-236 resource tags [\#54](https://github.com/rightscale/terraform-provider-rightscale/pull/54) ([mastamark](https://github.com/mastamark))
- STEEL-278 rescue retries and cloud polling [\#53](https://github.com/rightscale/terraform-provider-rightscale/pull/53) ([mastamark](https://github.com/mastamark))
- STEEL-109 Add more rsc tests [\#52](https://github.com/rightscale/terraform-provider-rightscale/pull/52) ([crunis](https://github.com/crunis))
- STEEL-276 Fix SecurityGroup doc [\#51](https://github.com/rightscale/terraform-provider-rightscale/pull/51) ([crunis](https://github.com/crunis))
- STEEL-270 custom provision functions [\#50](https://github.com/rightscale/terraform-provider-rightscale/pull/50) ([mastamark](https://github.com/mastamark))

## [v1.1.0](https://github.com/rightscale/terraform-provider-rightscale/tree/v1.1.0) (2018-05-18)
[Full Changelog](https://github.com/rightscale/terraform-provider-rightscale/compare/v1.0.0...v1.1.0)

**Merged pull requests:**

- STEEL-262 Ensure acceptance tests concurrency [\#49](https://github.com/rightscale/terraform-provider-rightscale/pull/49) ([crunis](https://github.com/crunis))
- Steel 239 server resource needs inputs to be useful [\#48](https://github.com/rightscale/terraform-provider-rightscale/pull/48) ([mastamark](https://github.com/mastamark))
- STEEL-261 - fix website erb template with doc links and fix ssh\_key m… [\#46](https://github.com/rightscale/terraform-provider-rightscale/pull/46) ([mastamark](https://github.com/mastamark))
- STEEL-258 Fix Sec Group Rule Race Cond in Tests [\#45](https://github.com/rightscale/terraform-provider-rightscale/pull/45) ([crunis](https://github.com/crunis))

## [v1.0.0](https://github.com/rightscale/terraform-provider-rightscale/tree/v1.0.0) (2018-05-08)
[Full Changelog](https://github.com/rightscale/terraform-provider-rightscale/compare/v0.0.1-alpha...v1.0.0)

**Closed issues:**

- Need RightScale Credential object support [\#20](https://github.com/rightscale/terraform-provider-rightscale/issues/20)

**Merged pull requests:**

- STEEL-259 - final cleanups of docs and readme for 1.0 release [\#44](https://github.com/rightscale/terraform-provider-rightscale/pull/44) ([mastamark](https://github.com/mastamark))
- Steel 90 route [\#43](https://github.com/rightscale/terraform-provider-rightscale/pull/43) ([mastamark](https://github.com/mastamark))
- STEEL-76 server\_array resource [\#42](https://github.com/rightscale/terraform-provider-rightscale/pull/42) ([crunis](https://github.com/crunis))
- STEEL-91 route table [\#41](https://github.com/rightscale/terraform-provider-rightscale/pull/41) ([mastamark](https://github.com/mastamark))
- Steel 255 use vpc networks for tests so we can use newer instance types [\#40](https://github.com/rightscale/terraform-provider-rightscale/pull/40) ([mastamark](https://github.com/mastamark))
- Steel 89 subnet resource and datasource [\#39](https://github.com/rightscale/terraform-provider-rightscale/pull/39) ([mastamark](https://github.com/mastamark))
- Steel 78 security group rule [\#38](https://github.com/rightscale/terraform-provider-rightscale/pull/38) ([mastamark](https://github.com/mastamark))
- STEEL-77 - sg resource and datasource polished and tests written [\#37](https://github.com/rightscale/terraform-provider-rightscale/pull/37) ([mastamark](https://github.com/mastamark))
- STEEl-256 - fix defaulting behavior for deployment resource and tag s… [\#36](https://github.com/rightscale/terraform-provider-rightscale/pull/36) ([mastamark](https://github.com/mastamark))
- STEEl-92 - network gateway resource fixed up, tests written and docum… [\#35](https://github.com/rightscale/terraform-provider-rightscale/pull/35) ([mastamark](https://github.com/mastamark))
- Steel-88 network resource, docs and tests [\#34](https://github.com/rightscale/terraform-provider-rightscale/pull/34) ([mastamark](https://github.com/mastamark))
- STEEL-248 datacenter docs [\#33](https://github.com/rightscale/terraform-provider-rightscale/pull/33) ([crunis](https://github.com/crunis))
- STEEL-249 multi cloud image docs [\#32](https://github.com/rightscale/terraform-provider-rightscale/pull/32) ([crunis](https://github.com/crunis))
- STEEL-250 server template docs [\#31](https://github.com/rightscale/terraform-provider-rightscale/pull/31) ([crunis](https://github.com/crunis))
- STEEL-251 Volume Snapshot Documentation [\#30](https://github.com/rightscale/terraform-provider-rightscale/pull/30) ([crunis](https://github.com/crunis))
- STEEL-237 Instance documentation [\#29](https://github.com/rightscale/terraform-provider-rightscale/pull/29) ([crunis](https://github.com/crunis))
- STEEL-252 - volume\_type datasource verified and docs added [\#28](https://github.com/rightscale/terraform-provider-rightscale/pull/28) ([mastamark](https://github.com/mastamark))
- Steel 238 image and instance type datasource verify and document [\#27](https://github.com/rightscale/terraform-provider-rightscale/pull/27) ([mastamark](https://github.com/mastamark))
- STEEL-232 substitute travis\_wait [\#26](https://github.com/rightscale/terraform-provider-rightscale/pull/26) ([crunis](https://github.com/crunis))
- STEEL-227 deployment resource and datasource [\#24](https://github.com/rightscale/terraform-provider-rightscale/pull/24) ([mastamark](https://github.com/mastamark))
- STEEL-95 resource\_cwf\_process improvements [\#23](https://github.com/rightscale/terraform-provider-rightscale/pull/23) ([crunis](https://github.com/crunis))
- STEEL-171 Set up travis-ci.org config [\#21](https://github.com/rightscale/terraform-provider-rightscale/pull/21) ([crunis](https://github.com/crunis))

## [v0.0.1-alpha](https://github.com/rightscale/terraform-provider-rightscale/tree/v0.0.1-alpha) (2018-02-15)
**Merged pull requests:**

- Add credential support [\#19](https://github.com/rightscale/terraform-provider-rightscale/pull/19) ([mastamark](https://github.com/mastamark))
- STEEL-160 Add server data source [\#18](https://github.com/rightscale/terraform-provider-rightscale/pull/18) ([bill-rich](https://github.com/bill-rich))
- STEEL-79 Volume DataSource Documentation [\#17](https://github.com/rightscale/terraform-provider-rightscale/pull/17) ([crunis](https://github.com/crunis))
- Add alpha notice to readme [\#16](https://github.com/rightscale/terraform-provider-rightscale/pull/16) ([adamalex](https://github.com/adamalex))
- STEEL-149 Use random string for functional test [\#15](https://github.com/rightscale/terraform-provider-rightscale/pull/15) ([adamalex](https://github.com/adamalex))
- STEEL-75 Fix up server resource and add tests and docs [\#14](https://github.com/rightscale/terraform-provider-rightscale/pull/14) ([bill-rich](https://github.com/bill-rich))
- Steel 121 tf datasources implement views [\#12](https://github.com/rightscale/terraform-provider-rightscale/pull/12) ([mark-dotson-rs](https://github.com/mark-dotson-rs))
- STEEL-132 Add unpopulated field filter [\#11](https://github.com/rightscale/terraform-provider-rightscale/pull/11) ([bill-rich](https://github.com/bill-rich))
- Steel 74 review prep [\#10](https://github.com/rightscale/terraform-provider-rightscale/pull/10) ([adamalex](https://github.com/adamalex))
- Steel 84 terraform rs ssh key resource tests and docs [\#9](https://github.com/rightscale/terraform-provider-rightscale/pull/9) ([mark-dotson-rs](https://github.com/mark-dotson-rs))
- STEEL-120 - Modify cmFilters function to properly format filter strin… [\#8](https://github.com/rightscale/terraform-provider-rightscale/pull/8) ([mark-dotson-rs](https://github.com/mark-dotson-rs))
- fix output race condition [\#7](https://github.com/rightscale/terraform-provider-rightscale/pull/7) ([adamalex](https://github.com/adamalex))
- Steel 101 travis should run acceptance tests on pull request creation [\#6](https://github.com/rightscale/terraform-provider-rightscale/pull/6) ([mark-dotson-rs](https://github.com/mark-dotson-rs))
- Steel 96 mix in a dash of ci [\#4](https://github.com/rightscale/terraform-provider-rightscale/pull/4) ([mark-dotson-rs](https://github.com/mark-dotson-rs))
- STEEL-73 - get repo uusing dep for dependenency management [\#3](https://github.com/rightscale/terraform-provider-rightscale/pull/3) ([mark-dotson-rs](https://github.com/mark-dotson-rs))
- STEEL-62 Fix cm instance resource and add tests [\#2](https://github.com/rightscale/terraform-provider-rightscale/pull/2) ([bill-rich](https://github.com/bill-rich))
- STEEL-60 Improve error messages for rsc client [\#1](https://github.com/rightscale/terraform-provider-rightscale/pull/1) ([bill-rich](https://github.com/bill-rich))



\* *This Change Log was automatically generated by [github_changelog_generator](https://github.com/skywinder/Github-Changelog-Generator)*