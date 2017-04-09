#!/bin/bash

#urun.sh
#You can use this bash file if you wanna run it safely.

#Usage:
#
#	"urun" <command> <flags>.
#


printHelp () {
	echo "";
	echo "USO is a simple chatting website, and this file,";
	echo "urun.sh, is the booter of this website\'s server.";
	echo "";
	echo "Usage:";
	echo "        [./]urun.sh <command> <flags>";
	echo "";
	echo "If you wanna run it, type \`urun.sh start\`.";
	echo "";
};

build () {
	go build -o ./uso.out ./*.go;
};

#The following function runs the uso.out.
#Usage:
# { echo noise; }|run;
# { echo quite; }|run;
run () {
	read run_mode;
	echo $run_mode;
	if [[ $run_mode == quite ]]; then {
		echo $run_mode;
		echo "--quite";
		./uso.out 1>./u.std.log 2>./u.err.log;
	}; elif [[ $run_mode == noise ]]; then {
		echo "--noise";
		./uso.out;
	}; fi;
};

#The following function extracts the arguments of this bash,
#analize it, and ...
parseArgs () {
	if [[ $BASH_ARGC -eq 0 ]]; then {
		#$ urun.sh;
		printHelp;
	}; else {
		build;
		command=${BASH_ARGV[$(($BASH_ARGC-1))]};
		echo The command is $command;
		if [[ $command == start ]]; then {
			{ echo noise; }|run;
		}; elif [[ $command == qstart ]]; then {
			{ echo quite; }|run;
		}; else {
			echo $command is an invalid command!;
		};fi;
	}; fi;
};

parseArgs;
