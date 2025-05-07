'use strict';
const TodoController = require('./todoController');

module.exports = function (app, { redisClient, logChannel }) {
  const todoController = new TodoController({ redisClient, logChannel });

  app.route('/todos')
    .get((req, res) => todoController.list(req, res))
    .post((req, res) => todoController.create(req, res));

  app.route('/todos/:taskId')
    .delete((req, res) => todoController.delete(req, res));
};
