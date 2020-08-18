<div class="container">
    <div class="naruto">
        <div class="body"></div>
        <div class="neck"></div>
        <div class="face">
            <div class="eyes"></div>
            <div class="mouth"></div>
            <div class="whiskers w1"></div>
            <div class="whiskers w2"></div>
            <div class="dimples"></div>
        </div>
        <div class="hair-base">
            <div class="hair h1"></div>
            <div class="hair h2"></div>
            <div class="hair h3"></div>
        </div>
        <div class="headband">
            <div class="plate"></div>
        </div>
        <!--div.noodle-->
        <div class="noodleLive"></div>
        <div class="chopsticks"></div>
        <div class="hand left"></div>
        <div class="hand right"></div>
    </div>
</div>
<div class="bowl"></div>
<div class="love"></div>
<style>
    .container {
        margin: 40px auto;
        width: 200px;
        height: 200px;
        position: relative;
        background: #A6DFE3;
        border-radius: 50%;
        overflow: hidden;
        z-index: 100;
    }

    .container div {
        position: absolute;
    }

    .naruto {
        width: 70%;
        height: 70%;
        bottom: 0px;
        left: 15%;
    }

    .face {
        width: 70%;
        height: 65%;
        border-radius: 50%;
        left: 15%;
        background: #FBE1D5;
    }

    .face:before,
    .face:after {
        content: "";
        position: absolute;
        border-radius: 50%;
        width: 25px;
        height: 20px;
        z-index: -5;
        background: #f9d0be;
    }

    .face:before {
        top: 38px;
        left: -6px;
    }

    .face:after {
        top: 38px;
        left: 80px;
    }

    .face .eyes:before,
    .face .eyes:after {
        content: "";
        position: absolute;
        width: 25px;
        height: 20px;
        border: 2px solid transparent;
        border-radius: 50%;
        border-top-color: #6b7686;
        top: 40px;
        z-index: 5;
    }

    .face .eyes:before {
        left: 10px;
    }

    .face .eyes:after {
        left: 60px;
    }

    .face .mouth:before,
    .face .mouth:after {
        content: "";
        position: absolute;
        width: 18px;
        height: 18px;
        border: 2px solid transparent;
        border-radius: 50%;
        border-bottom-color: #6b7686;
        top: 46px;
        z-index: 5;
    }

    .face .mouth:before {
        left: 32px;
    }

    .face .mouth:after {
        left: 46px;
    }

    .face .whiskers,
    .face .whiskers:before,
    .face .whiskers:after {
        position: absolute;
        width: 1px;
        height: 15px;
        background: #6b7686;
    }

    .face .w1 {
        top: 65%;
        left: 14%;
        transform: rotate(60deg);
    }

    .face .w2 {
        top: 65%;
        left: 84%;
        transform: rotate(-60deg);
    }

    .face .whiskers:before,
    .face .whiskers:after {
        content: "";
    }

    .face .whiskers:before {
        height: 18px;
    }

    .face .whiskers:after {
        height: 10px;
    }

    .face .w1:before {
        transform: translate(-8px, -2px) rotate(2deg);
    }

    .face .w1:after {
        transform: translate(8px, 4px);
    }

    .face .w2:before {
        transform: translate(8px, -2px) rotate(2deg);
    }

    .face .w2:after {
        transform: translate(-8px, 4px);
    }

    .hair-base {
        width: 70%;
        height: 20%;
        background: #FED84C;
        border-radius: 50%;
        left: 15%;
        top: -10px;
    }

    .hair {
        border: 30px solid transparent;
        border-bottom-width: 60px;
        border-bottom-color: #FED84C;
    }

    .hair:before,
    .hair:after {
        content: "";
        position: absolute;
        border: 15px solid transparent;
        border-bottom-width: 30px;
        border-bottom-color: #FED84C;
    }

    .hair.h1 {
        top: -65px;
        left: -10px;
        transform: rotate(-15deg);
    }

    .hair.h1:before {
        top: 10px;
        left: -45px;
        transform: rotate(-40deg);
    }

    .hair.h1:after {
        top: -5px;
        left: 8px;
        transform: rotate(20deg);
    }

    .hair.h2 {
        top: -60px;
        left: 35px;
        transform: rotate(15deg);
    }

    .hair.h2:before {
        top: 10px;
        left: 25px;
        transform: rotate(40deg);
    }

    .hair.h2:after {
        top: -10px;
        left: 15px;
        transform: rotate(20deg);
    }

    .hair.h3,
    .hair.h3:before,
    .hair.h3:after {
        border: 5px solid transparent;
        border-bottom-width: 15px;
        border-bottom-color: #FED84C;
        transform: rotate(180deg);
        z-index: 30;
    }

    .hair.h3 {
        top: 15px;
        left: 15px;
    }

    .hair.h3:before {
        top: -3px;
        left: 5px;
        transform: rotate(10deg);
    }

    .hair.h3:after {
        top: -2px;
        left: -70px;
        transform: rotate(-10deg);
    }

    .headband {
        width: 70%;
        height: 18%;
        background: #4977AF;
        top: 8%;
        left: 15%;
    }

    .headband:before,
    .headband:after {
        content: "";
        position: absolute;
        width: 100%;
        height: 15px;
        border-radius: 50%;
    }

    .headband:before {
        background: #4977AF;
        top: -6px;
    }

    .headband:after {
        background: #FBE1D5;
        top: 18px;
    }

    .headband .plate {
        width: 40%;
        height: 70%;
        border-radius: 10px;
        background: #BFC3CD;
        top: -3px;
        left: 30%;
    }

    .noodle {
        width: 20px;
        height: 60px;
        background: linear-gradient(to right, #fecb5c 0%, #fecb5c 15%, transparent 15%, transparent 25%, #fecb5c 25%, #fecb5c 40%, transparent 40%, transparent 50%, #fecb5c 50%, #fecb5c 65%, transparent 65%, transparent 75%, #fecb5c 75%, #fecb5c 90%, transparent 90%, transparent 100%);
        top: 68px;
        left: 60px;
    }

    .noodleLive {
        width: 20px;
        height: 60px;
        background: url(https://s12.postimg.org/c5vcr15od/noodle.png);
        animation: noodleScroll 4s linear infinite;
        top: 66px;
        left: 61px;
    }

    .neck {
        width: 80px;
        height: 15px;
        top: 55%;
        left: 20%;
        border-radius: 6px;
        background: linear-gradient(to right, white 0%, white 20%, #6b7686 20%, #6b7686 21%, white 21%, white 40%, #6b7686 40%, #6b7686 41%, white 41%, white 60%, #6b7686 60%, #6b7686 61%, white 61%, white 80%, #6b7686 80%, #6b7686 81%, white 81%);
    }

    .body {
        border: 110px solid transparent;
        border-top-color: #FDA142;
        border-radius: 50%;
        left: -35px;
        top: 80px;
    }

    .chopsticks {
        top: 40px;
        left: 5px;
    }

    .chopsticks:before,
    .chopsticks:after {
        content: "";
        position: absolute;
        height: 60px;
        width: 5px;
        background: #6b7686;
    }

    .chopsticks:before {
        transform: translate(10px) rotate(-30deg);
    }

    .chopsticks:after {
        transform: rotate(-35deg);
    }

    .bowl {
        position: absolute;
        width: 120px;
        height: 120px;
        border-radius: 50%;
        background: linear-gradient(to bottom, transparent 0, transparent 50%, #9CB654 50%, #9CB654 60%, white 60%, white 61%, #9CB654 61%);
        z-index: 120;
        top: 135px;
        left: calc(50% - 60px);
    }

    .hand {
        width: 30px;
        height: 60px;
        background: #FDA142;
        border-radius: 40%;
        z-index: 130;
    }

    .hand:before {
        content: "";
        position: absolute;
        width: 25px;
        height: 25px;
        border-radius: 50%;
        background: #FBE1D5;
    }

    .hand.left {
        top: 75px;
        left: -10px;
        transform: rotate(30deg);
    }

    .hand.left:before {
        transform: translate(2px, -10px);
    }

    .hand.right {
        top: 100px;
        left: 90%;
        transform: rotate(-20deg);
    }

    .hand.right:before {
        transform: translate(2px, -10px);
    }

    .love {
        position: absolute;
        z-index: 130;
        left: 53%;
        top: 80px;
        animation: love1 2s ease-out infinite;
    }

    .love:before,
    .love:after {
        position: absolute;
        content: "";
        left: 50px;
        top: 0;
        width: 20px;
        height: 30px;
        background: #F37A6E;
        border-radius: 50px 50px 0 0;
        transform: rotate(-25deg);
    }

    .love:after {
        transform: translate(8px, 2px) rotate(65deg);
    }

    @keyframes noodleScroll {
        from {
            background-position: 0 0;
        }
        to {
            background-position: 0 -120px;
        }
    }

    @keyframes love1 {
        from {
            transform: scale(0.2);
        }
        to {
            transform: translate(0px, -30px) scale(1);
        }
    }
</style>
