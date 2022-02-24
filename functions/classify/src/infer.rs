use wasmedge_tensorflow_interface;
use std::time::{Instant};

pub fn infer_internal(image_data: &[u8]) -> String {
    
    // 加载训练好的 TensorFlow lite 模型。
    let model_data: &[u8] = include_bytes!("models/mobilenet_v2/mobilenet_v2_1.4_224_frozen.pb");
    let labels = include_str!("models/mobilenet_v2/imagenet_slim_labels.txt");

    // let img_buf = base64::decode_config(&image_data, base64::STANDARD).unwrap();

    let start = Instant::now();
    // RGB32f
    let img = image::load_from_memory(image_data).unwrap().to_rgb8();
    println!("Loaded image in ... {:?}", start.elapsed());
    let resized = image::imageops::thumbnail(&img, 224, 224);
    println!("Resized image in ... {:?}", start.elapsed());
    let mut flat_img: Vec<f32> = Vec::new();

    for rgb in resized.pixels() {
        flat_img.push(rgb[0] as f32 / 255.);
        flat_img.push(rgb[1] as f32 / 255.);
        flat_img.push(rgb[2] as f32 / 255.);
    }

    // 加载上传图像并将其调整为224x224，这是这个 MobileNet 模型所需的尺寸，后面会介绍如何快速获得这个数据
    // let flat_img = wasmedge_tensorflow_interface::load_jpg_image_to_rgb32f(&img, 224, 224);  代码会出错 [error] jpeg is invalid. 

    // 用图像作为输入张量运行模型，并获取模型输出张量。
    let mut session = wasmedge_tensorflow_interface::Session::new(
        &model_data,
        // wasmedge_tensorflow_interface::ModelType::TensorFlowLite,
        wasmedge_tensorflow_interface::ModelType::TensorFlow,
    );

    session
    .add_input("input", &flat_img, &[1, 224, 224, 3])
    .add_output("MobilenetV2/Predictions/Softmax")
    .run();

    let res_vec: Vec<f32> = session.get_output("MobilenetV2/Predictions/Softmax");
    let mut i = 0;
    let mut max_index: i32 = -1;
    let mut max_value: f32 = -1.0;
    while i < res_vec.len() {
        let cur = res_vec[i];
        if cur > max_value {
            max_value = cur;
            max_index = i as i32;
        }
        i += 1;
    }
    // println!("{} : {}", max_index, max_value as f32 / 255.0);

    let mut confidence_zh = "可能有";
    // let mut confidence_en = "could be";  200
    if max_value > 0.75 {
        confidence_zh = "非常可能有";
        // confidence_en = "is very likely"; 125
    } else if max_value > 0.5 {
        confidence_zh = "很可能有";
        // confidence_en = "is likely"; 50
    } else if max_value > 0.2 {
        confidence_zh = "可能有";
        // confidence_en = "could be";
    }
    let mut label_lines = labels.lines();
    for _i in 0..max_index {
        label_lines.next();
    }

    let class_name = label_lines.next().unwrap().to_string();

    if max_value != 0.0 {
        //     "It {} a <a href='https://www.google.com/search?q={}'>{}</a> in the picture.",
        // format!("上传的图片里面{} <a href='https://www.baidu.com/s?wd={}'>{}</a>。{}:{}", confidence_zh.to_string(), class_name, class_name, max_index, max_value as f64 / 255.0)
        format!("{{\"confidence\":\"{}\",\"baidu\":\"https://www.baidu.com/s?wd={}\",\"google\":\"https://www.google.com/search?q={}\",\"name\":\"{}\",\"max_index\":\"{}\",\"max_value\":\"{}\"}}", confidence_zh.to_string(), class_name, class_name, class_name, max_index, max_value as f64 / 255.0)
    } else {
        // format!("上传的图片里面没有检测到 max_index:{}, max_value:{}", max_index, max_value as f64 / 255.0)
        format!("{{\"info\":\"It does not appears to be any in the picture.\",\"max_index\":\"{}\",\"max_value\":\"{}\"}}", max_index, max_value as f64 / 255.0)
    }
}
