#!/bin/bash
ssh -i "./virginia.pem" ubuntu@ec2-54-174-46-65.compute-1.amazonaws.com "cd /home/ubuntu/lifesaver && pkill -f lifesaver-server"
scp -r -i "./virginia.pem" lifesaver-server ubuntu@ec2-54-174-46-65.compute-1.amazonaws.com:/home/ubuntu/lifesaver
ssh -i "./virginia.pem" ubuntu@ec2-54-174-46-65.compute-1.amazonaws.com "cd /home/ubuntu/lifesaver && md5sum lifesaver-server && chmod +x lifesaver-server"
ssh -i "./virginia.pem" ubuntu@ec2-54-174-46-65.compute-1.amazonaws.com "cd /home/ubuntu/lifesaver && sh start.sh"

