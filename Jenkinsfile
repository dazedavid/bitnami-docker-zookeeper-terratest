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
               }
            }
         }
      }   
   } 
}