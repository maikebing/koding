$ = require 'jquery'
assert = require 'assert'

#! 389d1e79-20da-4519-8f85-cb7b73e8ddd4
# title: create_team_with_a_new_user
# start_uri: /
# tags: embedded
#

describe "create_team_with_a_new_user.rfml", ->
  describe """Click on the 'create a new team' link below the form where it says 'Welcome! Enter your team's Koding domain.'""", ->
    before -> 
      # implement before hook 

    it """Do you see 'Let's sign you up!' form with 'Email Address' and 'Team Name' text fields and a 'NEXT' button??""", -> 
      assert(false, 'Not Implemented')
      #assertion here


  describe """Enter "{{ random.email }}" in the 'Email Address' and "{{ random.last_name }}{{ random.number }}" in 'Team Name' fields and click on 'NEXT' button""", ->
    before -> 
      # implement before hook 

    it """Do you see 'Your team URL' form with 'Your Team URL' field pre-filled??""", -> 
      assert(false, 'Not Implemented')
      #assertion here


  describe """Click on 'NEXT' button""", ->
    before -> 
      # implement before hook 

    it """Do you see 'Your account' form with 'Email address' field pre-filled?""", -> 
      assert(false, 'Not Implemented')
      #assertion here

    it """Do you see 'Your Username' and 'Your Password' text fields??""", -> 
      assert(false, 'Not Implemented')
      #assertion here


  describe """Enter "{{ random.first_name }}{{ random.number }}" in the 'Your Username' and "{{ random.password }}" in the 'Your Password' field and click on 'CREATE YOUR TEAM' button""", ->
    before -> 
      # implement before hook 

    it """Do you see 'Authentication Required' form?""", -> 
      assert(false, 'Not Implemented')
      #assertion here

    it """(If not it's OK, then are you signed up successfully?)?""", -> 
      assert(false, 'Not Implemented')
      #assertion here


  describe """If you see 'Authentication Required' form, enter "koding" in the 'User Name:' and "1q2w3e4r" in the 'Password:' fields and click on 'Log In', if not do nothing and check the items below""", ->
    before -> 
      # implement before hook 

    it """Do you see 'You are almost there, {{ random.first_name }}{{ random.number }}!' title?""", -> 
      assert(false, 'Not Implemented')
      #assertion here

    it """Do you see 'Create a Stack for Your Team', 'Enter your Credentials', 'Build Your Stack', 'Invite Your Team' and 'Install KD' sections?""", -> 
      assert(false, 'Not Implemented')
      #assertion here

    it """Do you see 'CREATE', 'ENTER', 'BUILD', 'INVITE' and 'INSTALL' links next to them??""", -> 
      assert(false, 'Not Implemented')
      #assertion here


