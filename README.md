# WBABEProject-10

## 프로젝트 이름
띵동주문이요, 온라인 주문 시스템(Online Ordering System)

## 프로젝트 배경
언택트 시대에 급증하고 있는 온라인 주문 시스템은 이미 생활전반에 그 영향을 끼치고 있는 상황에, 가깝게는 배달어플, 매장에는 키오스크, 식당에는 패드를 이용한 메뉴 주문까지 그 사용범위가 점점 확대되어 가고 있습니다. 이런 시대에 해당 시스템을 이해, 경험하고 각 단계별 프로세스를 이해하여 구현함으로써 서비스 구축에 경험을 쌓고, golang의 이해를 돕습니다.

1. 학습자는 주문자/피주문자의 역할에서 필수적인 기능을 도출, 구현합니다.
2. 학습자는 해당 시스템에 대해 요구사항을 접수하고 주문자와 피주문자 입장에서 필요한 기능을 도출하여, 기능을 확장하고 주문 서비스를 원할하게 지원할수 있는 기능을 구현합니다.
3. 주문자는 신뢰있는 주문과 배달까지를 원합니다. 또, 피주문자는 주문내역을 관리하고 원할한 서비스가 제공되어야 합니다.

## 프로젝트 실행
```
[setting]

git clone https://github.com/codestates/WBABEProject-10.git
cd /WBABEProject-10
go mod tidy

[run]

go run main.go
```

## 디렉토리 구조
<img width="258" alt="image" src="https://user-images.githubusercontent.com/80724255/209470181-04f9999f-9444-4db7-a80d-6bbcdffad2f4.png">

## 데이터베이스 구조

### 메뉴(Menu)
- 메뉴 이름(name): string
- 주문 가능 여부(can_be_order): bool
- 수량(quantity): int
- 가격(price): int
- 원산지(origin): string
- 금일 추천 메뉴(today_recommend): bool
- 생성일(created_at): date
- 수정일(updated_at): date
- 삭제 여부(is_deleted): bool
    
### 리뷰(Review)
- 메뉴 정보(Menu)
- 주문자 정보(Orderer)
- 점수(score): int
- 추천 여부(is_recommend): bool
- 리뷰(review): string
- 생성일(created_at): date
- 수정일(updated_at): date
- 삭제 여부(is_deleted): bool
    
## 주문(Order)
- 메뉴 정보(Menu)
- 주문자 정보(Orderer)
- 상태(state): int (접수 - 0/조리중 - 1/배달중 - 2/배달완료- 3/주문취소-4)
- 주문 번호(numbering): int
- 생성일(created_at): date
- 수정일(updated_at): date

## 주문자(Orderer)
- 이름(name): varchar
- 전화번호(phone): varchar
- 주소(address): varchar
- 생성일(created_at): date

# API 명세

## Swagger 명세 확인
- 프로젝트 실행 후 http://localhost:8080/swagger/index.html#/로 접속

<img width="958" alt="image" src="https://user-images.githubusercontent.com/80724255/209470515-b0a611f8-9667-4654-b6fd-b5d0acf683fa.png">

## 피주문자 API

### 메뉴 추가
- [POST] /receipient/v01/menus
<img width="847" alt="image" src="https://user-images.githubusercontent.com/80724255/209675596-3e5dd964-864e-4187-8b7a-863f345cc423.png">

### 메뉴 삭제
- [DELETE] /receipient/v01/menus/{name}
<img width="856" alt="image" src="https://user-images.githubusercontent.com/80724255/209677634-41fabf79-31bd-4d3c-b321-0ef4b734df8d.png">

### 메뉴 수정
- [PATCH] /receipient/v01/menus/{name}
<img width="843" alt="image" src="https://user-images.githubusercontent.com/80724255/209678020-a411d8dd-d89d-4699-803a-bcd43c20556c.png">

### 주문 상태 변경
- [PATCH] /receipient/v01/order/{id}/state
<img width="851" alt="image" src="https://user-images.githubusercontent.com/80724255/209676205-f74b2936-830e-45fe-ac86-1c965d5d927f.png">

### 주문 확인
- [GET] /receipient/v01/order
<img width="858" alt="image" src="https://user-images.githubusercontent.com/80724255/209676701-0718330c-f272-49f8-8aa7-75751d2f1576.png">

### 주문자 API

### 주문 기능
- [POST]
<img width="855" alt="image" src="https://user-images.githubusercontent.com/80724255/209678564-e95017b1-8717-40fa-8f48-637c6b53bfd0.png">

### 주문 추가
- [PATCH] /orderer/v01/order/{id}
<img width="851" alt="image" src="https://user-images.githubusercontent.com/80724255/209676552-e0eb823f-e7bd-4a3c-9973-2069d0a54fd9.png">

### 주문 수정
- [PUT] /orderer/v01/order/{id}
<img width="853" alt="image" src="https://user-images.githubusercontent.com/80724255/209676410-03f11f13-bb7e-498d-9cd4-73dd454a95ec.png">

### 메뉴 조회
- [GET] /orderer/v01/menus
<img width="854" alt="image" src="https://user-images.githubusercontent.com/80724255/209675860-7addc3da-6804-4c9c-b41f-02ed18a6c6a9.png">

### 주문 상태 확인
- [GET] /orderer/v01/order/state?phone={phone}&address={address}
<img width="844" alt="image" src="https://user-images.githubusercontent.com/80724255/209675993-cbf24676-58d3-45da-bbee-9d2a90bedab9.png">

### 리뷰 생성
- [POST] /orderer/v01/reviews/{orderId}
<img width="854" alt="image" src="https://user-images.githubusercontent.com/80724255/209676862-746bcd3b-63fa-4da0-9482-0fae9b23fc69.png">

### 리뷰 조회
- [GET] /orderer/v01/reviews
<img width="853" alt="image" src="https://user-images.githubusercontent.com/80724255/209677106-58c3c26d-c52b-4246-9e6c-c769721cc256.png">

