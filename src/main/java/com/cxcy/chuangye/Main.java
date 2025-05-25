package com.cxcy.chuangye;

import com.alibaba.dashscope.audio.asr.recognition.Recognition;
import com.alibaba.dashscope.audio.asr.recognition.RecognitionParam;
import java.io.File;

public class Main {
    public static void main(String[] args) {
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
            System.out.println("识别结果：" + recognizer.call(param, new File("asr_example.wav")));
        } catch (Exception e) {
            e.printStackTrace();
        }
        System.out.println(
                "[Metric] requestId: "
                        + recognizer.getLastRequestId()
                        + ", first package delay ms: "
                        + recognizer.getFirstPackageDelay()
                        + ", last package delay ms: "
                        + recognizer.getLastPackageDelay());
        System.exit(0);
    }
}