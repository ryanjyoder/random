#!/usr/bin/env sh


echo hi setting the dns for rpi.local

# This script will continuously ping 'devhost' and if it's successful it will created a record in /etc/hosts that posts 'remotehost' to the ip address of 'devhost'
remotehost=archive.ubuntu.com
devhost=rpi.local

hostsfile=/etc/hosts
timeout=1

# we start not knowning the ip of the dev box
devip="unknown"
devip_update(){
	newip=$1
	if [ "$devip" = "$newip" ]; then
		return
	fi
	if [ "$newip" = "" ]; then
		# no ip was found unset the host
		unset_host_ip $remotehost
	else
		set_new_host_ip $remotehost $newip
	fi 
	devip=$newip
}

unset_host_ip(){
	newhostname=$1
	echo unsetting $newhostname
	tempfilename=`tempfile -m 0644`
	grep -v " $newhostname$" $hostsfile > $tempfilename
	cat $tempfilename > $hostsfile
	rm $tempfilename
}

set_new_host_ip(){
	newhostname=$1
	newhostip=$2
	echo setting new ip \($newhostip\) for $newhostname
	tempfilename=`tempfile -m 0644`
	grep -v " $newhostname$" $hostsfile > $tempfilename
	echo $newhostip $newhostname >> $tempfilename
	cat $tempfilename > $hostsfile
	rm $tempfilename
}

while ( true ); do
	curip=`avahi-resolve -n $devhost 2>/dev/null | sed 's/^.*\t\(.*\)$/\1/'`
	devip_update $curip 
	sleep $timeout
done

