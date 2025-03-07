# GitHub 用户活动监控工具

这是一个用于监控 GitHub 用户活动的命令行工具，支持多用户配置和实时活动查看。

## 功能特点

- 支持多用户配置管理
- 交互式用户选择
- 实时监控用户活动
- 显示活动类型、仓库和时间信息

## 安装

1. 克隆仓库：

```bash
git clone https://github.com/60tonAngel/github_user_activity.git
cd github_user_activity
```

2. 安装依赖：

```bash
go mod tidy
```

3. 编译项目：

```bash
go build
```

## 配置

在项目根目录下创建 `config` 文件夹，并在其中创建 `config.yaml` 文件：

```yaml
users:
  username1:    # GitHub 用户名
    token: "your_github_token"
  username2:
    token: "another_github_token"
```

注意：
- 确保 `token` 具有正确的访问权限
- YAML 文件的缩进必须正确
- 可以配置多个用户

## 使用方法

### 交互式选择用户

```bash
./github_user_activity select-user
```

这将显示所有配置的用户列表，让你选择要监控的用户。

### 直接监控指定用户

```bash
./github_user_activity monitor <username>
```

注意：指定的用户必须已在配置文件中配置。

## 输出示例

```
用户 username 的最近活动:
- 类型: PushEvent, 仓库: username/repo, 时间: 2024-03-07T10:30:00Z
- 类型: IssueCommentEvent, 仓库: org/project, 时间: 2024-03-07T09:15:00Z
```

## 错误处理

常见错误及解决方案：

1. "错误: 没有配置任何用户"
   - 检查 config.yaml 文件是否存在
   - 确认配置文件格式正确

2. "错误: 用户 xxx 未配置"
   - 确保指定的用户在配置文件中已配置

3. "获取用户活动失败"
   - 检查 GitHub token 是否有效
   - 确认网络连接正常

## 开发

要添加新功能或修改现有功能：

1. Fork 项目
2. 创建新分支
3. 提交更改
4. 发起 Pull Request
