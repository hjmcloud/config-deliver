# configDeliver

配置分发服务端

## 下载地址

- [linux_amd64](http://hjmcloud.gitee.io/config-deliver-release/linux_amd64/config-deliver)
- [linux_386](http://hjmcloud.gitee.io/config-deliver-release/linux_386/config-deliver)
- [darwin_amd64](http://hjmcloud.gitee.io/config-deliver-release/darwin_amd64/config-deliver)

## 配置文件

```yaml
# config.yaml 位于应用同级目录
server:
  address: ":8000"

logger:
  level: "info"
  stdout: true

dirmap:
  5472E3F6-37DE-B9A0-899C-39838E8C1336: ./test
```

## systemd 部署

本部署方式仅支持 linux 系统,且需要安装 systemd.

### 在线安装

进入要安装到的目录，执行以下命令

```shell
curl -sSL https://gitee.com/hjmcloud/scripts/raw/master/config-deliver/systemd/config-deliver-install.sh | sh
```

### 在线卸载

```shell
curl -sSL https://gitee.com/hjmcloud/scripts/raw/master/config-deliver/systemd/config-deliver-uninstall.sh | sh
```

### 服务管理

```shell
# 启动服务
systemctl start config-deliver
# 停止服务
systemctl stop config-deliver
# 重启服务
systemctl restart config-deliver
# 查看服务状态
systemctl status config-deliver
# 查看服务日志
journalctl -u config-deliver -f
```

## 开发命令

Project Makefile Commands:

- `make cli`: Install or Update to the latest GoFrame CLI tool.
- `make dao`: Generate go files for `Entity/DAO/DO` according to the configuration file from `hack` folder.
- `make service`: Parse `logic` folder to generate interface go files into `service` folder.
- `make image TAG=xxx`: Run `docker build` to build image according `manifest/docker`.
- `make image.push TAG=xxx`: Run `docker build` and `docker push` to build and push image according `manifest/docker`.
- `make deploy TAG=xxx`: Run `kustomize build` to build and deploy deployment to kubernetes server group according `manifest/deploy`.
- `make bin.publish`: build and publish binary files to gitee pages.
