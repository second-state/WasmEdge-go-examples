use wasmedge_bindgen::*;
use wasmedge_bindgen_macro::*;
use ndarray::{Array2};
use std::str::FromStr;

#[wasmedge_bindgen]
pub fn fit (csv_content: Vec<u8>, dim: i32, num_clusters: i32) -> String {
    let data = read_data(&csv_content, dim as usize);
    let (means, _clusters) = rkm::kmeans_lloyd(&data.view(), num_clusters as usize);

    // The following code groups the points into clusters around the means 
    // let data_view = data.view();
    // let groups = separate_groups(&data_view, &clusters);
    return serde_json::to_string(&means).unwrap();
}

fn read_data(csv_content: &[u8], dim: usize) -> Array2<f32> {
    let v : Vec<u8> = csv_content.to_vec();
    println!("INPUT length is {}", v.len());

    let mut data_reader = csv::Reader::from_reader(csv_content);
    let mut data: Vec<f32> = Vec::new();
    for record in data_reader.records() {
        for field in record.unwrap().iter() {
            let value = f32::from_str(field);
            data.push(value.unwrap());
        }
    }
    Array2::from_shape_vec((data.len() / dim, dim), data).unwrap()
}
