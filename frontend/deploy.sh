
#!/usr/bin/env sh

# 确保脚本抛出遇到的错误
set -e

# 生成静态文件
yarn run build

# 进入生成的文件夹
cd dist

# 如果发布到自定义域名
#echo 'b.hfbhfb.com' > CNAME
# 如果发布到 https://<username>.github.io

# 如果发布到 https://<username>.github.io/<repo>
git init
git add -A
git commit -m "deploy"
git push -f https://github.com/hfbhfb/qushuiyin.git master:gh-pages
