use wasmedge_tensorflow_interface;
use wasmedge_bindgen::*;
use wasmedge_bindgen_macro::*;

#[wasmedge_bindgen]
fn infer(image_data: Vec<u8>) -> String {
    let img = match image::load_from_memory(&image_data[..]) {
        Ok(a) => a.to_rgb8(),
        Err(e) => {
            println!("{:?}", e);
            panic!();
        }
    };
    let flat_img = image::imageops::thumbnail(&img, 192, 192);

    let model_data: &[u8] = include_bytes!("lite-model_aiy_vision_classifier_food_V1_1.tflite");
    let labels = include_str!("aiy_food_V1_labelmap.txt");

    let mut session = wasmedge_tensorflow_interface::Session::new(
        model_data,
        wasmedge_tensorflow_interface::ModelType::TensorFlowLite,
    );
    session
        .add_input("input", &flat_img, &[1, 192, 192, 3])
        .run();
    let res_vec: Vec<u8> = session.get_output("MobilenetV1/Predictions/Softmax");

    let mut i = 0;
    let mut max_index: i32 = -1;
    let mut max_value: u8 = 0;
    while i < res_vec.len() {
        let cur = res_vec[i];
        if cur > max_value {
            max_value = cur;
            max_index = i as i32;
        }
        i += 1;
    }

    let mut confidence = "could be";
    if max_value > 200 {
        confidence = "is very likely";
    } else if max_value > 125 {
        confidence = "is likely";
    } else if max_value > 50 {
        confidence = "could be";
    }

    let mut label_lines = labels.lines();
    for _i in 0..max_index {
      label_lines.next();
    }

    let class_name = label_lines.next().unwrap().to_string();
    if max_value > 50 {
      return format!("It {} a {} in the picture", confidence.to_string(), class_name);
    } else {
      return "It does not appears to be any food item in the picture.".to_string();
    }
}
