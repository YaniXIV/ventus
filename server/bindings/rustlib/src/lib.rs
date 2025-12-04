use acir_field::FieldElement;
use acir_field::AcirField;
use bn254_blackbox_solver::poseidon_hash;


fn bytes_to_field(bytes: &[u8]) -> FieldElement {
    FieldElement::from_be_bytes_reduce(bytes)
}


#[unsafe(no_mangle)]
pub unsafe extern "C" fn PoseidonHashGo(
    data1: *const u8,
    len1: usize,
    data2: *const u8,
    len2: usize,
    out_ptr: *mut u8,   // pointer to Go-allocated buffer
    out_len: *mut usize, // return actual length
) {
    let slice1 = std::slice::from_raw_parts(data1, len1);
    let left_fe = bytes_to_field(slice1);

    let result_fe = if !data2.is_null() && len2 > 0 {
        let slice2 = std::slice::from_raw_parts(data2, len2);
        let right_fe = bytes_to_field(slice2);
        poseidon_hash(&[left_fe, right_fe], false).unwrap()
    } else {
        poseidon_hash(&[left_fe], false).unwrap()
    };

    let out_bytes = result_fe.to_be_bytes(); // [u8; 32]

    // ---- copy into Go-allocated buffer ----
    if !out_ptr.is_null() {
        std::ptr::copy_nonoverlapping(out_bytes.as_ptr(), out_ptr, out_bytes.len());
    }

    if !out_len.is_null() {
        *out_len = out_bytes.len();
    }
}

