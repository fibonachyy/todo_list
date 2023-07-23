// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract TodoList {
    struct Task {
        uint256 id;
        string content;
        bool isCompleted;
    }

    mapping(address => Task[]) private todoLists;
    mapping(address => uint256) private taskCounter;

    event TaskAdded(address indexed user, uint256 taskId, string content);
    event TaskCompleted(address indexed user, uint256 taskId);

    function addTask(string memory _content) public {
        require(bytes(_content).length > 0, "Content should not be empty.");

        Task memory newTask = Task(taskCounter[msg.sender], _content, false);
        todoLists[msg.sender].push(newTask);
        taskCounter[msg.sender]++;

        emit TaskAdded(msg.sender, newTask.id, _content);
    }

    function markTaskCompleted(uint256 _taskId) public {
        require(_taskId < taskCounter[msg.sender], "Invalid task ID.");

        Task storage taskToUpdate = todoLists[msg.sender][_taskId];
        require(!taskToUpdate.isCompleted, "Task already completed.");

        taskToUpdate.isCompleted = true;
        emit TaskCompleted(msg.sender, _taskId);
    }

    function getTasksCount() public view returns (uint256) {
        return taskCounter[msg.sender];
    }

    function getTask(uint256 _taskId) public view returns (uint256, string memory, bool) {
        require(_taskId < taskCounter[msg.sender], "Invalid task ID.");

        Task storage task = todoLists[msg.sender][_taskId];
        return (task.id, task.content, task.isCompleted);
    }
}