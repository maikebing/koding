$                  = require 'jquery'
kd                 = require 'kd'
utils              = require './utils'

# ModalView = require 'app/integration-tests/modal'
# OutputModal = require 'app/integration-tests/output'
# Modal = require 'lab/Modal'


status =
  STARTED: 'started'
  RUNNING: 'running'
  PASSED: 'passed'
  FAILED: 'failed'


module.exports = class TestController extends kd.Controller
  constructor: ->
    super
    console.info 'Integration test initialized.'


  start: ->
    @emit 'status', status.STARTED

    @_run()


  # Bind to mocha test suite events, and emit when necessary.
  bindEvents: (runner) ->
    console.log 'runner' , runner
    runner.on 'end', =>
      @emit 'status', runner.currentRunnable.state


    runner.on 'fail', (res) =>
      console.log res, '>>>>>>>>> > > >> '
      { title, parent : _parent  } = res
      { message: status } = res.err
      { reactor } = kd.singletons
      title = "before hook of #{_parent.title}"  if title is '"before all" hook'

      # reactor.dispatch 'TEST_SUITE_FAIL', { title, status, parentTitle: @getParentTitle _parent}


  getParentTitle: (parent) ->

    return null if parent.title is ''
    while parent.parent.title isnt ''
      parent = parent.parent

    return parent.title


  # Setups mocha and add necessary tests.
  _run: () ->
    console.log '>>> mocha', mocha
    mocha.setup
      ui: 'bdd'
      timeout: 2000

    mocha.traceIgnores = [
      'https://cdnjs.cloudflare.com/ajax/libs/mocha/3.1.2/mocha.min.js'
    ]

    require './xxx'

    @emit 'status', status.RUNNING

    runner = mocha.run()

    @bindEvents runner


  prepareModal: ->

    # require 'app/integration-tests/style.css'

    # @modal = new OutputModal
    #   title: 'Testing Koding'
    #   isOpen: yes

    # @on 'status', (status) =>
    #   @modal.updateOptions
    #     title : "Testing Koding: #{status}"
    #     isOpen : yes

    @start()


  prepare: ->

    utils.loadMochaScript()

    resultView = new kd.CustomHTMLView
      domId: 'mocha'

    browserModal = new kd.ModalView
      cssClass : 'test-modal'
      title : 'Test Suites'
      overlay : yes
      buttons :
        Run :
          title : 'Run Tests'
          cssClass : 'kd solid medium'
          callback : =>
            @prepareModal()
        Cancel :
          title : 'Cancel'
          cssClass : 'kd kdbutton solid cancel medium'
          callback: -> browserModal.destroy()

    browserModal.addSubView resultView
