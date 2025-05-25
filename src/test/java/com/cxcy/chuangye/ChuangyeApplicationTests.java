package com.cxcy.chuangye;

import com.alibaba.dashscope.audio.asr.recognition.Recognition;
import com.alibaba.dashscope.audio.asr.recognition.RecognitionParam;
import com.cxcy.chuangye.service.AiService;
import jakarta.annotation.Resource;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;

import java.io.File;
import java.nio.charset.StandardCharsets;

@SpringBootTest
class ChuangyeApplicationTests {

    @Test
    void contextLoads() {

    }
    @Resource
    private AiService aiService;
    @Test
    public void callASRServiceTest() throws Exception {
        //读取文件
        File file = new File("C:\\Users\\niann\\Documents\\trae-projects\\boot-demo\\chuangye\\temp\\test.wav");
        //获取字节数组
        byte[] audioData = new byte[(int) file.length()];
        try {
            //读取文件
            java.io.FileInputStream fis = new java.io.FileInputStream(file);
            fis.read(audioData);
            fis.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
        String result = aiService.recognize(audioData, "wav");
        System.out.println(result);
    }


}
