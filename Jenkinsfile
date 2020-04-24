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
               sh 'cd test'
               sh 'go mod init "github.com/gruntwork-io/terratest/tree/master/modules"'  
               sh 'go test -v -tags kubernetes -run TestKubernetes'
               }
            }
         }
      }   
   } 
}
