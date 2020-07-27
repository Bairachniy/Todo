node {
    def root = tool name: 'go1.14', type: 'go'
        ws("${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/src/github.com/Bairachniy/Todo") {
        withEnv(["GOROOT=${root}", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/", "PATH+GO=${root}/bin"]) {
            env.PATH="${GOPATH}/bin:$PATH"
    
            git url: 'https://github.com/Bairachniy/Todo.git'
    
            stage 'preTest'
            sh 'go version'
            sh 'go get -u github.com/golang/dep/...'
            sh 'dep init'
            
            stage 'Test'
            sh 'go vet'
            sh 'go test -cover'
            
            stage 'Build'
            sh 'go build .'
            
            stage 'Deploy'
            // Do nothing.
}       
}
}
