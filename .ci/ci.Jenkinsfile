// Helper Functions
def checkoutCode(Map args=[:]){
  stage('Clone repository') {
    // CHECKOUT CODE REPO
    if(args.checkout_param){
      scmVars = checkout([
          $class: 'GitSCM',
          branches: [[name: "${env.SELECTED_BRANCH}"]],
          doGenerateSubmoduleConfigurations: false,
          extensions: [[$class: 'CleanCheckout'],
          [$class: 'PruneStaleBranch']],
          submoduleCfg: []
      ])
      // Re-init variables
      preparedVars()
    } else {
      scmVars = checkout scm
    }
  }
}

node(){
    checkoutCode()
    echo "Hello"
    currentBuild.result = 'FAILURE'
}