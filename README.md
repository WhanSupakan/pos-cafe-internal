# ☕ Cafe POS System

ระบบ Point of Sale (POS) สำหรับร้านกาแฟที่ออกแบบมาเพื่อแก้ไขปัญหาหลักๆ ที่พบในระบบ POS ทั่วไป โดยเน้นการจัดการสต็อกที่แม่นยำ การรายงานผลประกอบการที่ครบถ้วน การวิเคราะห์ประสิทธิภาพการตลาด และระบบจัดการสมาชิกลูกค้าที่เชื่อมโยงกับข้อมูลพฤติกรรม

## 📋 Problem

ระบบ POS ที่มีอยู่ในปัจจุบันยังมีข้อบกพร่องหลายประการที่ส่งผลกระทบต่อการดำเนินงานของร้านกาแฟ:

### 1. การนับ Stock ที่ไม่แม่นยำ
- **ปัญหาการนับสต็อกแบบ Manual**: ระบบส่วนใหญ่ยังต้องพึ่งพาการนับสต็อกด้วยมือ ซึ่งใช้เวลานานและเกิดข้อผิดพลาดได้ง่าย
- **ขาดการติดตาม Real-time**: ไม่สามารถทราบสถานะสต็อกแบบเรียลไทม์ ทำให้เกิดปัญหาสินค้าหมดหรือสต็อกค้าง
- **การเชื่อมโยงระหว่างการขายและสต็อกไม่ชัดเจน**: ไม่มีการอัพเดทสต็อกอัตโนมัติเมื่อมีการขาย ทำให้ข้อมูลไม่ตรงกับความเป็นจริง
- **ขาดการแจ้งเตือน**: ไม่มีการแจ้งเตือนเมื่อสต็อกใกล้หมดหรือหมด ทำให้พลาดโอกาสในการขาย

### 2. การสรุปผลรายได้ของร้านที่ไม่ครบถ้วน
- **รายงานที่จำกัด**: รายงานส่วนใหญ่แสดงเฉพาะยอดขายรวม ไม่มีการแยกตามประเภทสินค้า ช่วงเวลา หรือช่องทางการขาย
- **ขาดการวิเคราะห์แนวโน้ม**: ไม่สามารถวิเคราะห์แนวโน้มการขายได้ เช่น สินค้าขายดีในช่วงเวลาใด หรือวันไหนขายดีที่สุด
- **การคำนวณต้นทุนไม่แม่นยำ**: ไม่มีการคำนวณต้นทุนจริง (COGS) ที่เชื่อมโยงกับสต็อก ทำให้ไม่ทราบกำไรที่แท้จริง
- **ขาดการเปรียบเทียบ**: ไม่สามารถเปรียบเทียบผลประกอบการระหว่างช่วงเวลาได้อย่างง่ายดาย

### 3. การประเมินผล Marketing ที่ไม่มีประสิทธิภาพ
- **ไม่สามารถติดตาม ROI ของแคมเปญ**: ไม่มีระบบที่เชื่อมโยงระหว่างแคมเปญการตลาดกับยอดขายจริง
- **ขาดการวิเคราะห์ Customer Behavior**: ไม่ทราบว่าลูกค้าซื้อสินค้าอะไรบ้าง มาซื้อบ่อยแค่ไหน หรือมีค่าเฉลี่ยการซื้อเท่าไหร่
- **ไม่สามารถวัดประสิทธิภาพของโปรโมชัน**: ไม่ทราบว่าโปรโมชันใดที่ดึงดูดลูกค้าได้จริงและสร้างรายได้เพิ่ม
- **ขาดข้อมูลสำหรับการตัดสินใจ**: ไม่มีข้อมูลเพียงพอสำหรับการวางแผนแคมเปญการตลาดในอนาคต

### 4. การจัดการระบบสมาชิกลูกค้าที่ไม่ครบถ้วน
- **ขาดระบบสมาชิกที่เชื่อมโยงกับ Marketing**: ไม่มีระบบที่ต่อยอดจากข้อมูลการตลาดเพื่อสร้างโปรแกรมสมาชิกที่มีประสิทธิภาพ
- **ไม่สามารถติดตามพฤติกรรมลูกค้าได้ละเอียด**: ไม่ทราบว่าลูกค้าชอบซื้อสินค้าอะไร ทำอะไรบ่อย (เช่น มาเวลาไหน ชอบสั่งอะไร) หรือมีรูปแบบการซื้ออย่างไร
- **ระบบสะสมแต้มที่ไม่มีประสิทธิภาพ**: ระบบสะสมแต้มส่วนใหญ่เป็นแบบง่ายๆ ไม่มีการวิเคราะห์ว่าการให้แต้มสร้างผลตอบแทนจริงหรือไม่
- **ขาดการสร้าง Engagement**: ไม่มีระบบที่ช่วยสร้างความผูกพันกับลูกค้า เช่น การแจ้งเตือนโปรโมชันเฉพาะบุคคล หรือการให้รางวัลตามพฤติกรรม
- **ไม่สามารถ Segment ลูกค้าได้**: ไม่สามารถแบ่งกลุ่มลูกค้าตามพฤติกรรมเพื่อทำการตลาดแบบเจาะจง (Targeted Marketing)

## 💡 Solution

Cafe POS System ถูกออกแบบมาเพื่อแก้ไขปัญหาทั้งหมดข้างต้นด้วยแนวทางดังนี้:

### 1. ระบบจัดการสต็อกอัตโนมัติ
- **Real-time Stock Tracking**: อัพเดทสต็อกอัตโนมัติทุกครั้งที่มีการขายหรือรับสินค้าเข้า
- **การแจ้งเตือนอัตโนมัติ**: แจ้งเตือนเมื่อสต็อกใกล้หมดหรือหมด พร้อมคำแนะนำการสั่งซื้อ
- **การนับสต็อกที่แม่นยำ**: ระบบช่วยในการนับสต็อกจริง (Stock Take) และเปรียบเทียบกับข้อมูลในระบบ
- **การติดตามวัตถุดิบ**: ติดตามวัตถุดิบที่ใช้ในการผลิต (เช่น เมล็ดกาแฟ นม น้ำตาล) และคำนวณต้นทุนจริง

### 2. ระบบรายงานและวิเคราะห์ที่ครอบคลุม
- **Dashboard แบบ Real-time**: แสดงข้อมูลยอดขาย กำไร และสต็อกแบบเรียลไทม์
- **รายงานที่หลากหลาย**: 
  - รายงานยอดขายรายวัน/รายสัปดาห์/รายเดือน
  - รายงานสินค้าขายดี (Best Sellers)
  - รายงานกำไร-ขาดทุน (P&L Statement)
  - รายงานการเคลื่อนไหวสต็อก
- **การวิเคราะห์แนวโน้ม**: แสดงกราฟและเทรนด์การขายเพื่อช่วยในการวางแผน
- **การเปรียบเทียบช่วงเวลา**: เปรียบเทียบผลประกอบการระหว่างช่วงเวลาได้อย่างง่ายดาย

### 3. ระบบวิเคราะห์การตลาด
- **การติดตามแคมเปญ**: เชื่อมโยงแคมเปญการตลาดกับยอดขายจริง เพื่อคำนวณ ROI
- **Customer Analytics**: 
  - วิเคราะห์พฤติกรรมลูกค้า (Customer Behavior)
  - ติดตาม Customer Lifetime Value (CLV)
  - วิเคราะห์ความถี่ในการซื้อ (Purchase Frequency)
- **การวัดประสิทธิภาพโปรโมชัน**: ติดตามว่าโปรโมชันใดสร้างรายได้เพิ่มและดึงดูดลูกค้าได้จริง
- **ข้อมูลสำหรับการตัดสินใจ**: ให้ข้อมูลเชิงลึกสำหรับการวางแผนแคมเปญการตลาดในอนาคต

### 4. ระบบจัดการสมาชิกลูกค้าที่ครบถ้วน
- **ระบบสมาชิกที่ต่อยอดจาก Marketing**: 
  - เชื่อมโยงข้อมูลการตลาดกับโปรแกรมสมาชิก
  - วิเคราะห์ว่าลูกค้าคนไหนควรได้รับข้อเสนอพิเศษ
  - ติดตามประสิทธิภาพของโปรแกรมสมาชิก
- **การติดตามพฤติกรรมลูกค้าแบบละเอียด**:
  - บันทึกประวัติการซื้อทุกครั้ง (ซื้ออะไร เมื่อไหร่ ราคาเท่าไหร่)
  - วิเคราะห์รูปแบบการซื้อ (ซื้อบ่อยแค่ไหน ชอบซื้ออะไร ช่วงเวลาไหน)
  - สร้าง Customer Profile ที่ละเอียด
- **ระบบสะสมแต้มอัจฉริยะ**:
  - คำนวณแต้มอัตโนมัติตามยอดซื้อ
  - วิเคราะห์ ROI ของการให้แต้ม
  - ระบบแลกแต้มที่ยืดหยุ่น (แลกส่วนลด แลกสินค้า แลกสิทธิพิเศษ)
  - การหมดอายุแต้มและการแจ้งเตือน
- **การสร้าง Engagement**:
  - แจ้งเตือนโปรโมชันเฉพาะบุคคลตามพฤติกรรมการซื้อ
  - ระบบรางวัลตามความภักดี (Loyalty Rewards)
  - Birthday rewards และ special occasions
  - Push notifications สำหรับโปรโมชันที่เกี่ยวข้อง
- **Customer Segmentation**:
  - แบ่งกลุ่มลูกค้าอัตโนมัติตามพฤติกรรม (VIP, Regular, Occasional)
  - Targeted marketing สำหรับแต่ละกลุ่ม
  - วิเคราะห์ Customer Lifetime Value (CLV) เพื่อจัดลำดับความสำคัญ

## 🏗️ Architecture

โปรเจกต์นี้ใช้ **Clean Architecture** ร่วมกับ **MVC-style delivery** เพื่อให้:

- แยก business logic ออกจาก UI และ infrastructure
- รองรับการขยายตัวในอนาคต (เช่น เพิ่ม mobile app, API สำหรับ third-party)
- ทำให้โค้ดง่ายต่อการทดสอบและบำรุงรักษา
- ควบคุม dependency direction ให้ชัดเจน

### High-level Architecture

```
Frontend (React)
  → API Client
    → Backend HTTP Handler
      → Usecase (Business Logic)
        → Repository (Data Access)
          → Database/External Services
```

### Technology Stack

**Backend:**
- **Go** - ภาษาโปรแกรมหลัก
- **GORM** - ORM สำหรับ PostgreSQL
- **PostgreSQL** - ฐานข้อมูลหลัก
- **Redis** - Cache และ session management
- **Kafka** (optional) - Message queue สำหรับ async processing

**Frontend:**
- **React 18+** with **TypeScript**
- **React Router** - Routing
- **Axios** - HTTP client
- **Tailwind CSS** - Styling (recommended)

### Project Structure

```
cafe-pos/
├── backend/          # Go backend service
│   ├── internal/
│   │   ├── infrastructure/  # External dependencies (DB, Cache, API clients)
│   │   ├── config/          # Configuration
│   │   ├── db/models/       # Shared database models
│   │   ├── user/            # User module
│   │   │   ├── delivery/http/    # HTTP handlers
│   │   │   ├── handler/          # Controllers
│   │   │   ├── usecase/          # Business logic
│   │   │   └── repository/       # Data access
│   │   ├── order/           # Order module (same structure)
│   │   ├── customer/        # Customer/Membership module
│   │   │   ├── delivery/http/
│   │   │   ├── handler/
│   │   │   ├── usecase/
│   │   │   └── repository/
│   │   └── loyalty/         # Loyalty/Points module
│   │       ├── delivery/http/
│   │       ├── handler/
│   │       ├── usecase/
│   │       └── repository/
│   └── main.go
│
└── frontend/         # React frontend
    └── src/
        ├── infrastructure/   # API client, storage
        ├── config/           # Configuration
        ├── domain/models/    # Shared types
        ├── user/             # User module
        │   ├── presentation/ # UI components
        │   ├── application/  # Hooks, usecases, repositories
        │   └── domain/       # Module-specific types
        ├── order/            # Order module (same structure)
        ├── customer/         # Customer/Membership module
        │   ├── presentation/
        │   ├── application/
        │   └── domain/
        └── loyalty/          # Loyalty/Points module
            ├── presentation/
            ├── application/
            └── domain/
```

### Layer Responsibilities

- **Infrastructure**: External dependencies (database, cache, API clients)
- **Delivery/Presentation**: Transport layer (HTTP handlers, React components)
- **Handler/Hooks**: Orchestration layer (controllers, React hooks)
- **Usecase**: Business logic (pure functions, no framework dependencies)
- **Repository**: Data access (database queries, API calls)
- **Domain/Models**: Shared types and interfaces

สำหรับรายละเอียดเพิ่มเติม ดูได้ที่:
- [Backend Architecture](./backend/README.md)
- [Frontend Architecture](./frontend/README.md)

## 🎯 Key Decisions

### 1. Clean Architecture over Traditional MVC
**Decision**: ใช้ Clean Architecture แทน MVC แบบดั้งเดิม

**Rationale**:
- แยก business logic ออกจาก framework dependencies ทำให้ทดสอบง่ายขึ้น
- รองรับการเปลี่ยน technology stack ในอนาคตได้ง่าย
- ทำให้โค้ดมีโครงสร้างชัดเจนและง่ายต่อการบำรุงรักษา

**Trade-offs**:
- มีโค้ดมากกว่า MVC แบบดั้งเดิมเล็กน้อย
- ต้องเข้าใจ dependency rules

### 2. Module-based Structure
**Decision**: จัดโครงสร้างตาม module (user, order, product) แทนการจัดตาม layer

**Rationale**:
- ทำให้ง่ายต่อการค้นหาโค้ดที่เกี่ยวข้องกัน
- รองรับการทำงานแบบทีม (แต่ละทีมทำงานใน module ของตัวเอง)
- ลดความซับซ้อนเมื่อโปรเจกต์ขยายตัว

**Trade-offs**:
- อาจมีโค้ดซ้ำกันบ้างระหว่าง modules (แต่สามารถใช้ shared utilities ได้)

### 3. Real-time Stock Tracking
**Decision**: อัพเดทสต็อกแบบ real-time ทุกครั้งที่มีการขาย

**Rationale**:
- แก้ปัญหาการนับสต็อกที่ไม่แม่นยำ
- ให้ข้อมูลที่ถูกต้องสำหรับการตัดสินใจ
- ลดปัญหาสินค้าหมดโดยไม่ทราบ

**Trade-offs**:
- ต้องจัดการ concurrency เมื่อมีการขายพร้อมกันหลายรายการ
- อาจต้องใช้ database transaction เพื่อความถูกต้อง

### 4. Comprehensive Reporting System
**Decision**: สร้างระบบรายงานที่ครอบคลุมตั้งแต่แรก แทนการเพิ่มทีหลัง

**Rationale**:
- รายงานเป็น core feature ที่สำคัญสำหรับการจัดการร้าน
- ง่ายกว่าการเพิ่มทีหลังเพราะต้องเก็บข้อมูลให้ครบถ้วน
- ช่วยให้เจ้าของร้านตัดสินใจได้ดีขึ้น

**Trade-offs**:
- ใช้เวลาในการพัฒนามากขึ้น
- ต้องออกแบบ database schema ให้รองรับการ query ที่ซับซ้อน

### 5. Marketing Analytics Integration
**Decision**: รวมระบบวิเคราะห์การตลาดเข้าไปใน POS โดยตรง

**Rationale**:
- แก้ปัญหาการประเมินผล marketing ที่ไม่มีประสิทธิภาพ
- ให้ข้อมูลที่เชื่อมโยงกันระหว่างการขายและแคมเปญ
- ไม่ต้องพึ่งพาระบบภายนอก

**Trade-offs**:
- เพิ่มความซับซ้อนของระบบ
- ต้องออกแบบ data model ให้รองรับการติดตามแคมเปญ

### 6. Go for Backend
**Decision**: ใช้ Go แทน Node.js หรือ Python

**Rationale**:
- Performance สูง เหมาะกับระบบที่ต้องประมวลผลข้อมูลจำนวนมาก
- Concurrency model ที่ดี สำหรับจัดการ request พร้อมกัน
- Type safety ช่วยลด bugs
- Compile เป็น single binary ทำให้ deploy ง่าย

**Trade-offs**:
- Learning curve สำหรับทีมที่คุ้นเคยกับ JavaScript/Python
- Ecosystem น้อยกว่า Node.js/Python บ้าง

### 7. React with TypeScript for Frontend
**Decision**: ใช้ React + TypeScript แทน Vue หรือ vanilla JavaScript

**Rationale**:
- TypeScript ให้ type safety ช่วยลด runtime errors
- React ecosystem ใหญ่ มี library มากมาย
- Component-based architecture สอดคล้องกับ Clean Architecture
- รองรับการขยายเป็น React Native ในอนาคต

**Trade-offs**:
- ต้องเขียน type definitions เพิ่มเติม
- Bundle size ใหญ่กว่า vanilla JS

### 8. Integrated Customer Membership System
**Decision**: สร้างระบบสมาชิกที่เชื่อมโยงกับ Marketing และ Analytics แบบบูรณาการ

**Rationale**:
- ต่อยอดจากข้อมูลการตลาดเพื่อสร้างโปรแกรมสมาชิกที่มีประสิทธิภาพ
- การติดตามพฤติกรรมลูกค้าอย่างละเอียดช่วยให้สามารถทำ Targeted Marketing ได้
- ระบบสะสมแต้มที่เชื่อมโยงกับข้อมูลพฤติกรรมช่วยเพิ่ม Engagement
- Customer Segmentation ช่วยให้การตลาดมีประสิทธิภาพมากขึ้น

**Trade-offs**:
- ต้องเก็บข้อมูลลูกค้าเพิ่มเติม (privacy concerns)
- ระบบซับซ้อนขึ้น ต้องจัดการข้อมูลจำนวนมาก
- ต้องออกแบบ database schema ให้รองรับการ query ที่ซับซ้อน
- อาจต้องใช้ data analytics tools เพิ่มเติม

## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- Node.js 18+
- PostgreSQL 14+
- Redis 6+ (optional)

### Backend Setup
```bash
cd backend
go mod init cafe-pos
go mod tidy
# Configure .env file
go run main.go
```

### Frontend Setup
```bash
cd frontend
npm install
# Configure .env file
npm run dev
```

## 📝 License

[Specify your license here]

## 👥 Contributors

[Add contributors here]
