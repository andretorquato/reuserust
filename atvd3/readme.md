# Pattern Proxy

O **Design Pattern Proxy** é utilizado no Vue 3 para implementar reatividade, permitindo que valores sejam monitorados e atualizados automaticamente. Esse padrão é a base para duas funcionalidades principais do framework:

---

## Funcionalidades Reativas no Vue 3
 **`ref`:** 
   - Permite criar um objeto reativo encapsulando valores primitivos ou complexos.
   - Quando o valor encapsulado é alterado, as mudanças são automaticamente refletidas na interface do usuário.

 **`reactive`:**
   - Permite criar um objeto reativo a partir de objetos.
   - Toda alteração nas propriedades do objeto dispara atualizações automáticas na interface.
