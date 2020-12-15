if [ $1 = NO-CACHE ]
then
   docker build --no-cache --tag atlas-tds:latest .
else
   docker build --tag atlas-tds:latest .
fi
