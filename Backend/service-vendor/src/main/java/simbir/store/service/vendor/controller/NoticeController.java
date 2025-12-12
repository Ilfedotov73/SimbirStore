package simbir.store.service.vendor.controller;

import lombok.RequiredArgsConstructor;
import org.springframework.format.annotation.DateTimeFormat;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.client.RestTemplate;
import simbir.store.service.vendor.dto.NoticesDto;

import java.sql.Date;
import java.util.*;
import java.util.concurrent.CopyOnWriteArrayList;
import java.util.stream.Collectors;

@RestController
@RequestMapping("/api/notices")
@RequiredArgsConstructor
public class NoticeController {

    private final RestTemplate restTemplate = new RestTemplate();

    // URL другого микросервиса, который возвращает все уведомления
    private final String noticesServiceUrl = "http://notice-service/api/notices";

    // Локальный кэш для принятых уведомлений
    private final List<NoticesDto> localNotices = new CopyOnWriteArrayList<>();

    /**
     * Получить уведомления для конкретного продавца за период
     */
    @GetMapping("/vendor/{vendorId}")
    public List<NoticesDto> getVendorNotices(
            @PathVariable int vendorId,
            @RequestParam @DateTimeFormat(iso = DateTimeFormat.ISO.DATE) Date startDate,
            @RequestParam @DateTimeFormat(iso = DateTimeFormat.ISO.DATE) Date endDate
    ) {

        // Получаем уведомления с другого микросервиса
        NoticesDto[] externalNotices = restTemplate.getForObject(noticesServiceUrl, NoticesDto[].class);

        List<NoticesDto> allNotices = new ArrayList<>();
        if (externalNotices != null) {
            allNotices.addAll(Arrays.asList(externalNotices));
        }

        // Добавляем локально принятые уведомления
        allNotices.addAll(localNotices);

        // Фильтруем по vendorId и дате, сортируем по дате (последние первыми)
        return allNotices.stream()
                .filter(n -> n.getEntityId() == vendorId)
                .filter(n -> !n.getCreateAt().before(startDate) && !n.getCreateAt().after(endDate))
                .sorted((n1, n2) -> n2.getCreateAt().compareTo(n1.getCreateAt()))
                .toList();
    }

    /**
     * Приём нового уведомления от другого микросервиса
     */
    @PostMapping
    public NoticesDto receiveNotice(@RequestBody NoticesDto notice) {
        // Можно сохранять в базу, но для примера просто добавим в локальный кэш
        localNotices.add(notice);

        // Логируем или проксируем дальше, если нужно
        System.out.println("Received notice: " + notice.getText());

        return notice;
    }
}
