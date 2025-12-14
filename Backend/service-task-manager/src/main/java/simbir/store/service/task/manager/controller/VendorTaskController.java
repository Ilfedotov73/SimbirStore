package simbir.store.service.task.manager.controller;

import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import simbir.store.service.task.manager.db.entity.TaskVendorRegistration;
import simbir.store.service.task.manager.dto.TaskVendorRegistrationDto;
import simbir.store.service.task.manager.service.ServiceTaskManager;

import java.util.List;

@RestController
@RequestMapping("/api/v1")
@RequiredArgsConstructor
public class VendorTaskController {

    private  ServiceTaskManager taskManager;

    // ---------------------------------------------------------
    // GET /GetVenndorRegTasks
    // ---------------------------------------------------------
    @GetMapping("/GetVenndorRegTasks")
    public ResponseEntity<List<TaskVendorRegistration>> getVendorRegTasks() {
        return ResponseEntity.ok(
                taskManager.getVendorRegistrationTasks()
        );
    }

    // ---------------------------------------------------------
    // GET /GetVendorRegTaskById/{id}
    // ---------------------------------------------------------
    @GetMapping("/GetVendorRegTaskById/{id}")
    public ResponseEntity<TaskVendorRegistration> getVendorRegTaskById(
            @PathVariable Long id
    ) {
        return ResponseEntity.ok(
                taskManager.getVendorRegistrationTaskById(id)
        );
    }

    // ---------------------------------------------------------
    // POST /InsertVendorRegTask
    // ---------------------------------------------------------
    @PostMapping("/InsertVendorRegTask")
    public ResponseEntity<TaskVendorRegistrationDto> insertVendorRegTask(
            @RequestBody TaskVendorRegistrationDto dto
    ) {
        return ResponseEntity.ok(
                taskManager.createVendorRegistrationTask(dto)
        );
    }

    // ---------------------------------------------------------
    // PATCH /PatchVendorRegTask/{status}/{adminId}/{taskId}
    // ---------------------------------------------------------
    @PatchMapping("/PatchVendorRegTask/{status}/{adminId}/{taskId}")
    public ResponseEntity<Void> patchVendorRegTask(
            @PathVariable String status,
            @PathVariable Long adminId,
            @PathVariable Long taskId
    ) {
        taskManager.updateVendorRegistrationTaskStatus(taskId, adminId , status);
        return ResponseEntity.ok().build();
    }
}
