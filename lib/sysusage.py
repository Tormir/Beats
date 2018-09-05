#!/usr/bin/python3.6

import os, subprocess, json, argparse

def Parser():
    parser = argparse.ArgumentParser(description='Process metrix from system')
    parser.add_argument('-C', metavar='Command', type=str, help='Command which process started, delimiter is ":::", e.g. command:command', required=True)
    return parser.parse_args()

def Windows() :
    return None


def Linux() :
    parser = Parser()
    command = parser.C.split(":::")
    data = {}

    proces = subprocess.run(['ps', 'aux'], stdout=subprocess.PIPE).stdout.decode("utf-8").splitlines()
    for x in range(len(json.loads(json.dumps(proces)))):
        data[' '.join(json.loads(json.dumps(proces))[x].split()[10:])] = '{"USER":"' + json.loads(json.dumps(proces))[x].split()[0] + '"' + \
                                                                         ', "PID":' + json.loads(json.dumps(proces))[x].split()[1] + \
                                                                         ', "CPU":' + json.loads(json.dumps(proces))[x].split()[2] + \
                                                                         ', "MEM":' + json.loads(json.dumps(proces))[x].split()[3] + \
                                                                         ', "VSZ":' + json.loads(json.dumps(proces))[x].split()[4] + \
                                                                         ', "RSS":' + json.loads(json.dumps(proces))[x].split()[5] + \
                                                                         ', "TTY":"' + json.loads(json.dumps(proces))[x].split()[6] + '"' + \
                                                                         ', "STAT":"' + json.loads(json.dumps(proces))[x].split()[7] + '"' + \
                                                                         ', "START":"' + json.loads(json.dumps(proces))[x].split()[8] + '"' + \
                                                                         ', "TIME":"' + json.loads(json.dumps(proces))[x].split()[9] + '"' + \
                                                                         ', "COMMAND":"' + ' '.join(json.loads(json.dumps(proces))[x].split()[10:]) + '"' + \
                                                                         '}'

    with open('lib/metrics.json', 'w') as f:
        for i in command:
            f.write(data.get(i) + '\n')


if __name__ == '__main__':
    Linux()
