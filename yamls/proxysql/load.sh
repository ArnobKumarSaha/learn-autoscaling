COUNTER=0

USER='test'
PROXYSQL_NAME='my-proxy'
NAMESPACE='demo'
PASS='test'

VAR="x"

while [  $COUNTER -lt 40 ]; do
    let COUNTER=COUNTER+1
    VAR=a$VAR
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e 'select 1;' > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "INSERT INTO ok.okok(name) VALUES ('$VAR');" > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "select * from ok.okok;" > /dev/null 2>&1
    sleep 0.0001
done

VAR="x"
COUNTER=0
while [  $COUNTER -lt 40 ]; do
    let COUNTER=COUNTER+1
    VAR=b$VAR
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e 'select 1;' > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "INSERT INTO ok.okok(name) VALUES ('$VAR');" > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "select * from ok.okok;" > /dev/null 2>&1
    sleep 0.0001
done

VAR="x"
COUNTER=0
while [  $COUNTER -lt 40 ]; do
    let COUNTER=COUNTER+1
    VAR=c$VAR
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e 'select 1;' > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "INSERT INTO ok.okok(name) VALUES ('$VAR');" > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "select * from ok.okok;" > /dev/null 2>&1
    sleep 0.0001
done

VAR="x"
COUNTER=0
while [  $COUNTER -lt 40 ]; do
    let COUNTER=COUNTER+1
    VAR=d$VAR
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e 'select 1;' > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "INSERT INTO ok.okok(name) VALUES ('$VAR');" > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "select * from ok.okok;" > /dev/null 2>&1
    sleep 0.0001
done

VAR="x"
COUNTER=0
while [  $COUNTER -lt 40 ]; do
    let COUNTER=COUNTER+1
    VAR=e$VAR
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e 'select 1;' > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "INSERT INTO ok.okok(name) VALUES ('$VAR');" > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "select * from ok.okok;" > /dev/null 2>&1
    sleep 0.0001
done

VAR="x"
COUNTER=0
while [  $COUNTER -lt 40 ]; do
    let COUNTER=COUNTER+1
    VAR=f$VAR
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e 'select 1;' > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "INSERT INTO ok.okok(name) VALUES ('$VAR');" > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "select * from ok.okok;" > /dev/null 2>&1
    sleep 0.0001
done

VAR="x"
COUNTER=0
while [  $COUNTER -lt 40 ]; do
    let COUNTER=COUNTER+1
    VAR=g$VAR
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e 'select 1;' > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "INSERT INTO ok.okok(name) VALUES ('$VAR');" > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "select * from ok.okok;" > /dev/null 2>&1
    sleep 0.0001
done

VAR="x"
COUNTER=0
while [  $COUNTER -lt 40 ]; do
    let COUNTER=COUNTER+1
    VAR=h$VAR
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e 'select 1;' > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "INSERT INTO ok.okok(name) VALUES ('$VAR');" > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "select * from ok.okok;" > /dev/null 2>&1
    sleep 0.0001
done

VAR="x"
COUNTER=0
while [  $COUNTER -lt 40 ]; do
    let COUNTER=COUNTER+1
    VAR=i$VAR
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e 'select 1;' > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "INSERT INTO ok.okok(name) VALUES ('$VAR');" > /dev/null 2>&1
    mysql -u$USER -h$PROXYSQL_NAME.$NAMESPACE.svc -P6033 -p$PASS -e "select * from ok.okok;" > /dev/null 2>&1
    sleep 0.0001
done
