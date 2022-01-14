use forkscanner::{ForkScanner, Block, run_server};
use log::info;
use diesel::prelude::PgConnection;
use diesel::Connection;
use diesel::prelude::*;
use diesel::result::QueryResult;

fn main() {
    dotenv::dotenv().expect("Failed loading dotenv");
    let db_url = std::env::var("DATABASE_URL").expect("No DB url");
    let db_conn = PgConnection::establish(&db_url).expect("Connection failed");

    let scanner = ForkScanner::new(db_conn).expect("Launching forkscanner failed");
    let duration = std::time::Duration::from_millis(10_000);

    println!("Forkscanner is running...");

    let _handle_scanner = std::thread::spawn(move || {
        loop {
            scanner.run();
            info!("Run finished, sleeping");
            std::thread::sleep(duration);
        }
    });

    let _handle_kdeep_blocks = std::thread::spawn(move || {
        loop {
            get_kdeep_blocks();
            info!("Fetched k_deep blocks, sleeping");
            std::thread::sleep(duration);
        }
    });

    run_server("127.0.0.1:8339", &db_url);
}

pub fn get_kdeep_blocks_query(conn: &PgConnection) -> QueryResult<Vec<Block>> {
    let raw_query = format!("
        WITH RECURSIVE rec_query AS (
            SELECT * FROM blocks WHERE hash = (select block from chaintips where status = 'active' )
            UNION ALL
            SELECT b.* FROM blocks b INNER JOIN rec_query r ON r.parent_hash = b.hash
        ) SELECT * FROM rec_query ORDER BY height DESC Limit 10;
    ");

    diesel::sql_query(raw_query).load(conn)
}

fn get_kdeep_blocks() {
    dotenv::dotenv().expect("Failed loading dotenv");
    let db_url = std::env::var("DATABASE_URL").expect("No DB url");
    let db_conn = PgConnection::establish(&db_url).expect("Connection failed");
    let blocks = get_kdeep_blocks_query(&db_conn).unwrap();
    println!("{:?}", blocks[0].hash);
}