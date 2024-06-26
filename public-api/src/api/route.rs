use actix_web::{web, HttpResponse};
use utoipa::OpenApi;
use utoipa_swagger_ui::SwaggerUi;

use crate::api::openapi::ApiDoc;






async fn route_unmatch() -> Result<HttpResponse, actix_web::Error> {
    Ok(HttpResponse::NotFound().body("{\"message\": \"Not Found\"}"))
}

pub fn config(cfg: &mut web::ServiceConfig) {
    cfg.service(
        SwaggerUi::new("/docs/{_:.*}")
            .url("/v1/docs/openapi.json", ApiDoc::openapi())
    )
    .service(super::versions::versions)
    .service(super::earthquake::earthquake_eventids_hourly)
    .service(super::earthquake::earthquake_eventids_daily)
    .service(super::earthquake::get_earthquake_detail)
    .service(super::wether_warning::get_weather_warning)
    .service(super::public_transportation::get_public_transportation_area)
    .service(super::public_transportation::get_public_transportation_area_to_line)
    .service(super::public_transportation::get_public_transportation_line)
    .default_service(
        web::route().to(route_unmatch)
    );
}


