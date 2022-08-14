![](https://img.shields.io/badge/npm-7.13.0-blue)
![](https://img.shields.io/badge/node-10.19.0-blue)
![](https://img.shields.io/badge/vue/cli-4.5.13-brightgreen)
![](https://img.shields.io/badge/Go-1.15-orange)
![](https://img.shields.io/badge/Fiber-2.18.0-orange)

# Nicoreco-portal

## 環境構築

ローカルに clone

```
git clone https://github.com/kenkn/Nicoreco-portal.git
```

コンテナ起動

```
docker-compose up
```

### "vue-cli-service: command not found"の解決策

```
cd ui
rm -rf node_modules package-lock.json && npm install
```
