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
                 withKubeConfig([
                    credentialsId: 'jenkins-robot',
                    caCertificate: '',
                    serverUrl: 'https://172.31.33.66',
                    contextName: '',
                    clusterName: '',
                    namespace: ''
                    ]) {
               def root = tool name: 'Go'
               withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
               sh 'go version'
                  dir("test"){
                     sh 'pwd'
                     sh 'go test -v -tags kubernetes -run TestKubernetes'                  
                  }
               }
               }
            }
         }
      }   
   } 
}
