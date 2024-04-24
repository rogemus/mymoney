echo "Remove old build files..."
rm -rf ./tmp
rm -rf ./ui/public/browser


echo "Remove old build...\n"
sshpass -p $GH_PASS ssh $GH_USER@frog01.mikr.us -p 11893 "rm -f main-linux"

echo "Prepare .env"
printenv > .env

make build

echo "Uploading build files ...\n"
sshpass -p $GH_PASS scp -P 11893 ./tmp/main-linux $GH_USER@frog01.mikr.us:
