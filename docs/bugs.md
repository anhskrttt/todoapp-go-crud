# Bug Document

## Bug 01: Create a new post return weird response
* Status: Done
* Bug Point: ./controllers/controllers.go - CreateTask()
* Description: 
    * Tryna create a new task with the following data (Notice the comma in the end of the "description" line)
        ```
        {
            "title": "Task 01",
            "description": "This is a task test",
        }
        ```
    * Response: Warning Headers were already written. Wanted to override status code 400 with 200

* Solution: Use ShouldBind() instead of Bind()

## Bug 02: