pipeline {
   agent any
   stages {
      stage ('Checking out GIT Files') {
         steps {
            checkout scm
        }
      }
      stage ('Preparing Test') {
         steps {
            script {
               def root = tool name: 'Go'
               withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
               sh 'go version'
                  dir("test"){
                     sh 'pwd'
                     sh 'curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/darwin/amd64/kubectl"'
                     sh 'chmod +x ./kubectl'
                     sh 'ls -latr'
                     sh 'go mod init "github.com/gruntwork-io/terratest/tree/master/modules"'
                     sh 'go test -v -tags kubernetes -run TestKubernetes'                  
                  }
               }
            }
         }
      }   
   } 
}
