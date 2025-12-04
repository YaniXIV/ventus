use acir_field::FieldElement;
use bn254_blackbox_solver::poseidon_hash;



fn main() {
    let inputs = vec![
        FieldElement::from(1u128),
        FieldElement::from(2u128),
        FieldElement::from(3u128),
        FieldElement::from(4u128),
    ];
    let is_variable_length = false;
    let result = poseidon_hash(&inputs, is_variable_length).expect("should hash successfully");
    println!("Poseidon2 hash result: {:?}", result);
}
