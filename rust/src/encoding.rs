extern crate serialize;

use self::serialize::hex::FromHex;
use self::serialize::base64::{ToBase64, STANDARD};


pub fn hex_to_base64(hex: &str) -> String {
    hex.from_hex().unwrap().as_slice().to_base64(STANDARD)
}
