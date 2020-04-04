# gitea-vet

[![Build Status](https://drone.gitea.com/api/badges/jolheiser/gitea-vet/status.svg)](https://drone.gitea.com/jolheiser/gitea-vet)

`go vet` tool for Gitea

| Analyzer | Description                                                         |
|----------|---------------------------------------------------------------------|
| Imports  | Checks for import sorting. stdlib->code.gitea.io->other             |
| License  | Checks file headers for some form of `Copyright...YYYY...Gitea/Gogs`|
