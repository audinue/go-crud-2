<?php

return new class {
  private $auth = [
    'driver'   => 'pgsql',
    'server'   => 'localhost:5432',
    'username' => 'postgres',
    'password' => 'postgres',
    'db'       => 'postgres',
  ];

  function __construct()
  {
    if ($_SERVER['REQUEST_URI'] == '/') {
      $_POST['auth'] = $this->auth;
    }
  }

  function credentials()
  {
    return [
      $this->auth['server'],
      $this->auth['username'],
      $this->auth['password'],
    ];
  }

  function login()
  {
    return true;
  }
};
