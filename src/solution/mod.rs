mod heap;
mod linked_list;
mod s1;
mod s11;
mod s1109;
mod s12;
mod s121;
mod s122;
mod s125;
mod s128;
mod s13;
mod s134;
mod s135;
mod s14;
mod s1480;
mod s15;
mod s151;
mod s167;
mod s169;
mod s189;
mod s2;
mod s202;
mod s205;
mod s209;
mod s219;
mod s228;
mod s238;
mod s242;
mod s26;
mod s27;
mod s274;
mod s28;
mod s289;
mod s290;
mod s295;
mod s3;
mod s36;
mod s380;
mod s383;
mod s392;
mod s42;
mod s443;
mod s45;
mod s452;
mod s48;
mod s49;
mod s528;
mod s54;
mod s55;
mod s56;
mod s57;
mod s58;
mod s6;
mod s68;
mod s704;
mod s73;
mod s76;
mod s771;
mod s80;
mod s860;
mod s88;
mod s881;

#[macro_export]
macro_rules! vec_of_strings {
    // match a list of expressions separated by comma:
    ($($str:expr),*) => ({
        // create a Vec with this list of expressions,
        // calling String::from on each:
        vec![$(String::from($str),)*] as Vec<String>
    });
}
