# GitHub �û����ع���

����һ�����ڼ�� GitHub �û���������й��ߣ�֧�ֶ��û����ú�ʵʱ��鿴��

## �����ص�

- ֧�ֶ��û����ù���
- ����ʽ�û�ѡ��
- ʵʱ����û��
- ��ʾ����͡��ֿ��ʱ����Ϣ

## ��װ

1. ��¡�ֿ⣺

```bash
git clone https://github.com/60tonAngel/github_user_activity.git
cd github_user_activity
```

2. ��װ������

```bash
go mod tidy
```

3. ������Ŀ��

```bash
go build
```

## ����

����Ŀ��Ŀ¼�´��� `config` �ļ��У��������д��� `config.yaml` �ļ���

```yaml
users:
  username1:    # GitHub �û���
    token: "your_github_token"
  username2:
    token: "another_github_token"
```

ע�⣺
- ȷ�� `token` ������ȷ�ķ���Ȩ��
- YAML �ļ�������������ȷ
- �������ö���û�

## ʹ�÷���

### ����ʽѡ���û�

```bash
./github_user_activity select-user
```

�⽫��ʾ�������õ��û��б�����ѡ��Ҫ��ص��û���

### ֱ�Ӽ��ָ���û�

```bash
./github_user_activity monitor <username>
```

ע�⣺ָ�����û��������������ļ������á�

## ���ʾ��

```
�û� username ������:
- ����: PushEvent, �ֿ�: username/repo, ʱ��: 2024-03-07T10:30:00Z
- ����: IssueCommentEvent, �ֿ�: org/project, ʱ��: 2024-03-07T09:15:00Z
```

## ������

�������󼰽��������

1. "����: û�������κ��û�"
   - ��� config.yaml �ļ��Ƿ����
   - ȷ�������ļ���ʽ��ȷ

2. "����: �û� xxx δ����"
   - ȷ��ָ�����û��������ļ���������

3. "��ȡ�û��ʧ��"
   - ��� GitHub token �Ƿ���Ч
   - ȷ��������������

## ����

Ҫ����¹��ܻ��޸����й��ܣ�

1. Fork ��Ŀ
2. �����·�֧
3. �ύ����
4. ���� Pull Request
