#[allow(dead_code)]
#[derive(Default)]
struct RandomizedSet {
    list: Vec<i32>,
    map: std::collections::HashMap<i32, usize>,
}

#[allow(dead_code)]
impl RandomizedSet {
    fn new() -> Self {
        Default::default()
    }

    fn insert(&mut self, val: i32) -> bool {
        match self.map.get(&val) {
            Some(_) => false,
            None => {
                self.list.push(val);
                self.map.insert(val, self.list.len() - 1);
                true
            }
        }
    }

    fn remove(&mut self, val: i32) -> bool {
        match self.map.get(&val) {
            Some(&index) => {
                self.list.swap_remove(index);
                self.map.remove(&val);
                match self.list.get(index) {
                    Some(&v) => {
                        self.map.insert(v, index);
                    }
                    None => {}
                }
                true
            }
            None => false,
        }
    }

    fn get_random(&self) -> i32 {
        let mills = std::time::SystemTime::now()
            .duration_since(std::time::UNIX_EPOCH)
            .unwrap()
            .as_nanos() as i32;
        let index = mills as usize % self.list.len();
        self.list[index]
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_solution() {
        let mut set = RandomizedSet::new();
        assert_eq!(true, set.insert(1));
        assert_eq!(false, set.remove(2));
        assert_eq!(true, set.insert(2));
        // assert_eq!(1, set.get_random());
        assert_eq!(true, set.remove(1));
        assert_eq!(false, set.insert(2));
        assert_eq!(2, set.get_random());
    }
}
