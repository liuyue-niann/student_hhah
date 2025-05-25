package com.cxcy.chuangye.service;

import com.alibaba.dashscope.audio.asr.recognition.Recognition;
import com.alibaba.dashscope.audio.asr.recognition.RecognitionParam;
import com.alibaba.dashscope.audio.asr.recognition.RecognitionResult;
import com.alibaba.dashscope.common.ResultCallback;
import com.cxcy.chuangye.domain.dto.Root;
import com.google.gson.Gson;
import com.google.gson.JsonArray;
import com.google.gson.JsonObject;
import com.google.gson.JsonParser;
import org.springframework.stereotype.Service;

import javax.sound.sampled.AudioFormat;
import javax.sound.sampled.AudioInputStream;
import java.io.ByteArrayInputStream;
import java.io.File;
import java.io.FileOutputStream;
import java.nio.ByteBuffer;
import java.nio.charset.StandardCharsets;
import java.util.Arrays;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;

@Service
public class AiService {

    private static final String API_KEY = "sk-f9aff3cd3aae4fc9881d95188f8afd28";
    private static final int TIMEOUT_SECONDS = 30;

    /**
     * 识别音频数据
     * @param audioData 音频字节数组
     * @param format 音频格式（如"wav"）
     * @return 识别结果文本
     * @throws Exception 识别过程中可能出现的异常
     */
    public String recognize(byte[] audioData, String format) throws Exception {
        File file = new File("temp/./test.wav");
        try (FileOutputStream fos = new FileOutputStream(file)) {
            fos.write(audioData);
        }
        // 创建Recognition实例
        Recognition recognizer = new Recognition();
        // 创建RecognitionParam
        RecognitionParam param =
                RecognitionParam.builder()
                        // 若没有将API Key配置到环境变量中，需将下面这行代码注释放开，并将apiKey替换为自己的API Key
                        .apiKey("sk-f9aff3cd3aae4fc9881d95188f8afd28")
                        .model("paraformer-realtime-v2")
                        .format("wav")
                        .sampleRate(16000)
                        // “language_hints”只支持paraformer-realtime-v2模型
                        .parameter("language_hints", new String[]{"zh", "en"})
                        .build();

        try {
            String call = recognizer.call(param, file);
            System.out.println("result:"+call);
            Gson gson = new Gson();
            Root root = gson.fromJson(call, Root.class);
            if  (root.getSentences().isEmpty()){
                return "";
            }
            return root.getSentences().get(0).getText();
        } catch (Exception e) {
//            删除临时文件
            e.printStackTrace();
        }finally {
            file.delete();
        }
        System.out.println(
                "[Metric] requestId: "
                        + recognizer.getLastRequestId()
                        + ", first package delay ms: "
                        + recognizer.getFirstPackageDelay()
                        + ", last package delay ms: "
                        + recognizer.getLastPackageDelay());
        return "";
    }
}