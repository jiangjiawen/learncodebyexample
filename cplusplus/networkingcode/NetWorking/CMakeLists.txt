cmake_minimum_required(VERSION 3.16)
project(NetWorking)

set(CMAKE_CXX_STANDARD 17)

#find_library(CARBON_LIBRARY Carbon.framework)

set(ASIO_CMAKE_ASIO_TAG asio-1-18-1)
add_subdirectory(thirdparty/asio)

include_directories(NetCommon)

add_executable(SimpleClient NetClient/SimpleClient.cpp)
target_link_libraries(SimpleClient PRIVATE asio "-framework Carbon")

add_executable(SimpleServer NetServer/SimpleServer.cpp)
target_link_libraries(SimpleServer PRIVATE asio)