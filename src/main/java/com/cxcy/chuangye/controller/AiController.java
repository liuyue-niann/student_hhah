package com.cxcy.chuangye.controller;

import com.cxcy.chuangye.service.AiService;
import jakarta.annotation.Resource;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import javax.sound.sampled.*;
import java.io.*;
import java.util.HashMap;
import java.util.Map;

@CrossOrigin
@RestController
public class AiController {
    @Resource
    private AiService aiService;


    @PostMapping("/aiVoiceChat")
    public Map<String, Object> recognizeVoice(
            @RequestParam("audio") MultipartFile audio,
            @RequestParam(value = "format", required = false, defaultValue = "wav") String format) {
        Map<String, Object> result = new HashMap<>();

        File file = new File("./temp/test.wav");
        try (FileOutputStream fos = new FileOutputStream(file)) {
            fos.write(audio.getBytes());
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
        try {
            // TODO: 调用语音识别服务
            String recognize = aiService.recognize(audio.getBytes(), format);
            result.put("code", 200);
            result.put("message", "识别成功");
            result.put("data", Map.of("text", recognize));
        } catch (Exception e) {
            e.printStackTrace();
            result.put("code", 500);
            result.put("message", "识别失败: " + e.getMessage());
        }
        return result;
    }

}
