use anyhow::{Error, Result};
use log::debug;
use once_cell::sync::Lazy;
use public_api_lib::scheme::earthquake::EarthQuake;

/// Asynchronously collects earthquake data.
///
/// # Examples
///
/// ```rust
/// # use anyhow::Result;
/// # async fn example() -> Result<()> {
///     earthquake_operator().await;
///     Ok(())
/// # }
/// ```
#[allow(clippy::unused_unit)]
pub async fn earthquake_operator() {

    let earthquakes_data: Vec<EarthQuake>;

    match collect_earthquake().await {
        Ok(earthquakes) => {
            earthquakes_data = earthquakes;


            match public_api_lib::earthquake_operator::earthquake_data_submitter(&earthquakes_data).await {
                Ok(_) => {
                    log::info!("Earthquake Data Submitted Successfully");
                },
                Err(e) => {
                    log::error!("Error: {}", e);
                    return ();
                }
            }


        },
        Err(e) => {
            log::error!("Error: {}", e);
            return ()
        }
        
    }
    debug!("Earthquake Data: {:?}", earthquakes_data);
    



}

/// Collects earthquake data from the API.
///
/// # Returns
///
/// A vector of EarthQuake structs representing the collected earthquake data.
///
/// # Errors
///
/// Returns an error if there was a problem collecting the data or parsing the response.
///
/// # Examples
///
/// ```rust
/// # use anyhow::Result;
/// # async fn example() -> Result<()> {
///     let earthquakes = collect_earthquake().await?;
///     Ok(())
/// # }
/// ```
pub async fn collect_earthquake() -> Result<Vec<public_api_lib::scheme::earthquake::EarthQuake>, Error> {

    static BASE_URL: Lazy<std::string::String> = Lazy::new(|| {
        format!("{}/get/earthquake?max=50", *public_api_lib::get_env::API_URL)
    });

    log::info!("Collecting Earthquake Data from: {}", &BASE_URL.to_string());

    let response = reqwest::get(BASE_URL.as_str())
        .await?
        .text()
        .await?;    

    debug!("Recieved Response: {}", response);

    let deserialized : Vec<public_api_lib::scheme::earthquake::EarthQuake> = serde_json::from_str(&response)?;

    debug!("Converted: {:?}", deserialized);


    Ok(deserialized)
}





