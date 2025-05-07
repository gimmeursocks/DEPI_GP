'use strict';
const cache = require('memory-cache');

const OPERATION_CREATE = 'CREATE',
      OPERATION_DELETE = 'DELETE';

class TodoController {
    constructor({ redisClient, logChannel }) {
        this._redisClient = redisClient;
        this._logChannel = logChannel;
    }

    list(req, res) {
        const data = this._getTodoData(req.user.username);
        res.json(data.items);
    }

    create(req, res) {
        const data = this._getTodoData(req.user.username);
        const todo = {
            content: req.body.content,
            id: data.lastInsertedID
        };
        data.items[data.lastInsertedID] = todo;
        data.lastInsertedID++;
        this._setTodoData(req.user.username, data);

        this._logOperation(OPERATION_CREATE, req.user.username, todo.id);
        res.json(todo);
    }

    delete(req, res) {
        const data = this._getTodoData(req.user.username);
        const id = req.params.taskId;
        delete data.items[id];
        this._setTodoData(req.user.username, data);

        this._logOperation(OPERATION_DELETE, req.user.username, id);
        res.status(204).send();
    }

    _logOperation(opName, username, todoId) {
        this._redisClient.publish(this._logChannel, JSON.stringify({
            opName,
            username,
            todoId,
        }));
    }

    _getTodoData(userID) {
        let data = cache.get(userID);
        if (!data) {
            data = {
                items: {
                    '1': { id: 1, content: "Create new todo" },
                    '2': { id: 2, content: "Update me" },
                    '3': { id: 3, content: "Delete example ones" }
                },
                lastInsertedID: 4
            };
            this._setTodoData(userID, data);
        }
        return data;
    }

    _setTodoData(userID, data) {
        cache.put(userID, data);
    }
}

module.exports = TodoController;
