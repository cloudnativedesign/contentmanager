# CloudNativeDesign - Contentmanager
Manages the lifecycle of content on the platform across Articles, Social Media content, Videos, Blogs and linked media. 

The contentmanager service provides a detailed CRUD API for all content elements and provides optimized storage and life cycle management by integrating with the storage service. 

Content is registered across the platform for ease of retrievability and lifecycle events are created and persisted to the monitoring service. 

## Architecture
The service is decomposed into:
* Webserver
* REST API
* Controllers
* Models
* Database

And it integrates with the following services:
* StorageManager
* Monitor
* AnalyBrain

## Models
### Owned models
* Article
* Tweet
* BlogPost
* LinkedinStory
* Slideshare
* YoutubeVideo
* YoutubeChannel

### Mirrored models
* Author
* Person

## Component Roadmap
The service lays the foundational capacity to efficiently handle large amounts of registered content across the platform including its manual and automatic ingestion pipelines. It is envisoned to grow from a simple but comprehensive CRUD API to a smart, data enhacing content platform that can provide access to large distributed content sources efficiently and cost effectively. 

Many requirements for efficient data handling will be pushed down towards the storage service to handle the optimal management of optionally multiple storage backends also across locations. 

Furthermore, the registration and management of data streams will be made available with the abbility to manage metadata across running streams in near real-time for indexing, searchability, data join and data augmentation use cases across the platform. 
