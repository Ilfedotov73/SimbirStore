package simbir.store.service.task.manager.controller;

import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import simbir.store.service.task.manager.db.entity.TaskProductRegistration;
import simbir.store.service.task.manager.dto.TaskProductsRegistrationDto;
import simbir.store.service.task.manager.service.ServiceTaskManager;

import java.util.List;

@RestController
@RequestMapping("/api/v1")
@RequiredArgsConstructor
public class ProductTaskController {


    private  ServiceTaskManager taskManager;

    // ---------------------------------------------------------
    // GET /GetProductsRegTasks
    // ---------------------------------------------------------
    @GetMapping("/GetProductsRegTasks")
    public ResponseEntity<List<TaskProductRegistration>> getProductsRegTasks() {
        return ResponseEntity.ok(
                taskManager.getProductsRegistrationTasks()
        );
    }

    // ---------------------------------------------------------
    // GET /GetProductRegTaskById/{id}
    // ---------------------------------------------------------
    @GetMapping("/GetProductRegTaskById/{id}")
    public ResponseEntity<TaskProductRegistration> getProductRegTaskById(
            @PathVariable Long id
    ) {
        return ResponseEntity.ok(
                taskManager.getProductsRegistrationTaskById(id)
        );
    }

    // ---------------------------------------------------------
    // POST /InsertProductsRegTask
    // ---------------------------------------------------------
    @PostMapping("/InsertProductsRegTask")
    public ResponseEntity<TaskProductsRegistrationDto> insertProductsRegTask(
            @RequestBody TaskProductsRegistrationDto dto
    ) {
        return ResponseEntity.ok(
                taskManager.createProductsRegistrationTask(dto)
        );
    }

    // ---------------------------------------------------------
    // PATCH /PatchProductsRegTask/{status}/{adminId}/{taskId}
    // ---------------------------------------------------------
    @PatchMapping("/PatchProductsRegTask/{status}/{adminId}/{taskId}")
    public ResponseEntity<Void> patchProductRegTask(
            @PathVariable String status,
            @PathVariable Long adminId,
            @PathVariable Long taskId
    ) {
        taskManager.updateProductsRegistrationTaskStatus(taskId, adminId, status);
        return ResponseEntity.ok().build();
    }
}
