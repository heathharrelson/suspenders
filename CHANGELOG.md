# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Support running behind a proxy or ingress with new `--external-url` argument

## [0.2.0] - 2020-09-05

### Added

- Initial dashboard page with list of deployments
- Multi-arch Docker image: [heathharrelson/suspenders](https://hub.docker.com/r/heathharrelson/suspenders)
- Example Kubernetes manifest for cluster-scoped installation
- Health endpoint for readiness / liveness checks

## 0.1.0 - Never Released

[Unreleased]: https://github.com/heathharrelson/suspenders/compare/v0.2.0...Head
[0.2.0]: https://github.com/heathharrelson/suspenders/releases/tag/v0.2.0