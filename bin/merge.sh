#/bin/bash -
#
# 服务启动/停止
#

PID=0

usage() {
	echo "Usage: ./merge.sh [ -h ] [ -k start | stop | restart ]";
}

check() {
	PID=`ps aux | grep "./bin/merge " | grep -v "grep" | awk '{print $2}'`
}

start() {
	check;
	if [ -z "$PID" ]; then
		./bin/merge --maxIdleConns 3000 --maxIdleConnsPerHost 1000 >> /var/log/merge-requests.log 2>&1 &
	else
		echo "Already is running, pid:$PID"
	fi
}

stop() {
	check;
	if [ ! -z "$PID" ]; then
		#kill -SIGTERM $PID
		kill -9 $PID
	fi
}

restart() {
	stop;
	start;
}

cd $(dirname $0)
cd ../

while getopts k:h arg
do
	case $arg in
		k)
			ACTION=$OPTARG;;
		h)
			usage;
			exit;
	esac
done

if [ -$ACTION = -start ] || [ -$ACTION = -restart ] || [ -$ACTION = -stop ]
then
	$ACTION
else
	usage
fi
